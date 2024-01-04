package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	auth "github.com/ugabiga/falcon/internal/authentication"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"github.com/ugabiga/falcon/pkg/config"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

const (
	maxAge        = 86400 * 30
	callbackURL   = "/auth/signin/google/callback"
	JWTCookieName = "falcon.access_token"
)

type JWTClaim struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type WhiteListType string

const (
	WhiteListTypePrefix WhiteListType = "prefix"
	WhiteListTypeExact  WhiteListType = "exact"
)

type WhiteList struct {
	Type WhiteListType
	Path string
}

type AuthenticationService struct {
	db                 *ent.Client
	cfg                *config.Config
	authenticationRepo *repository.AuthenticationDynamoRepository
	userRepo           *repository.UserDynamoRepository
}

func NewAuthenticationService(
	db *ent.Client,
	cfg *config.Config,
	authenticationRepo *repository.AuthenticationDynamoRepository,
	userRepo *repository.UserDynamoRepository,

) *AuthenticationService {
	a := &AuthenticationService{
		db:                 db,
		cfg:                cfg,
		authenticationRepo: authenticationRepo,
		userRepo:           userRepo,
	}

	return a
}

func (s AuthenticationService) JWTMiddleware(whiteList []WhiteList) echo.MiddlewareFunc {
	secretKey := s.cfg.JWTSecretKey

	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaim)
		},
		SigningKey:  []byte(secretKey),
		TokenLookup: "header:Authorization:Bearer ,cookie:" + JWTCookieName,
		Skipper: func(c echo.Context) bool {
			for _, v := range whiteList {
				requestedPath := c.Path()

				switch v.Type {
				case WhiteListTypePrefix:
					if strings.HasPrefix(requestedPath, v.Path) {
						return true
					}
				case WhiteListTypeExact:
					if requestedPath == v.Path {
						return true
					}
				default:
					return false
				}
			}

			return false
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, map[string]string{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		},
	})
}

func (s AuthenticationService) UngradedJWTMiddleware() echo.MiddlewareFunc {
	secretKey := s.cfg.JWTSecretKey

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (returnErr error) {
			cookies := c.Cookies()
			if len(cookies) == 0 {
				return next(c)
			}

			tokenStr, err := c.Request().Cookie(JWTCookieName)
			if err != nil {
				return next(c)
			}

			token, err := jwt.ParseWithClaims(
				tokenStr.Value,
				&JWTClaim{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, err
					}
					return []byte(secretKey), nil
				},
			)
			if err != nil {
				return next(c)
			}

			c.Set("user", token)

			return next(c)
		}
	}

}

func (s AuthenticationService) CreateJWTToken(userID string, name string, isAdmin bool) (string, error) {
	secretKey := s.cfg.JWTSecretKey
	claims := &JWTClaim{
		userID,
		name,
		isAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s AuthenticationService) VerifyUser(ctx context.Context, loginType, token string) (*auth.UserProvidedData, error) {
	p, err := s.provider(loginType)
	if err != nil {
		return nil, err
	}

	user, err := p.GetUserData(ctx, &oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  token,
		RefreshToken: "",
	})

	return user, err
}

func (s AuthenticationService) SignInOrSignUp(
	ctx context.Context, provider string, identifier string, credential string, name string,
) (
	*model.Authentication, *model.User, error,
) {
	a, err := s.authenticationRepo.GetItem(ctx, provider, identifier)
	u, err := s.userRepo.Get(ctx, a.UserID)
	if err != nil {
		//TODO check if error is not found
		return s.SignUp(ctx, provider, identifier, credential, name)
	}

	return a, u, nil
}

func (s AuthenticationService) SignUp(
	ctx context.Context, authenticationProvider string, identifier string, credential string, name string,
) (
	*model.Authentication, *model.User, error,
) {
	defaultTimezone := "Asia/Seoul"

	inputUser := model.User{
		Timezone: defaultTimezone,
	}
	if name != "" {
		inputUser.Name = name
	}
	createdUser, err := s.userRepo.Create(ctx, inputUser)
	if err != nil {
		return nil, nil, err
	}

	inputAuthentication := model.Authentication{
		Identifier: identifier,
		Provider:   authenticationProvider,
		Credential: credential,
		UserID:     createdUser.ID,
	}

	createdAuthentication, err := s.authenticationRepo.Create(ctx, inputAuthentication)
	if err != nil {
		return nil, nil, err
	}

	return createdAuthentication, createdUser, nil
}

func (s AuthenticationService) provider(loginType string) (auth.OAuthProvider, error) {
	switch loginType {
	case "google":
		return auth.NewGoogleProvider(s.cfg.GoogleClientID, s.cfg.GoogleClientSecret)
	default:
		return nil, errors.New("unknown login type")
	}
}
