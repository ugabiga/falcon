package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/authentication"
	"strings"
	"time"
)

const (
	maxAge        = 86400 * 30
	baseURL       = "http://localhost:3000"
	callbackURL   = baseURL + "/auth/signin/google/callback"
	JWTCookieName = "falcon.access_token"
)

type JWTClaim struct {
	UserID  int    `json:"user_id"`
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
	db  *ent.Client
	cfg *config.Config
}

func NewAuthenticationService(db *ent.Client, cfg *config.Config) *AuthenticationService {
	a := &AuthenticationService{
		db:  db,
		cfg: cfg,
	}
	a.InitializeOAuthProviders()

	return a
}

func (s AuthenticationService) InitializeOAuthProviders() {
	secretKey := s.cfg.SessionSecretKey
	googleClientID := s.cfg.GoogleClientID
	googleClientSecret := s.cfg.GoogleClientSecret

	store := sessions.NewCookieStore([]byte(secretKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   false,
	}

	gothic.Store = store
	goth.UseProviders(
		google.New(
			googleClientID,
			googleClientSecret,
			callbackURL,
			"email", "profile",
		),
	)
}

func (s AuthenticationService) CreateJWTToken(userID int, name string, isAdmin bool) (string, error) {
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

func (s AuthenticationService) JWTClaim(c echo.Context) (*JWTClaim, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, echo.ErrUnauthorized
	}
	claims := user.Claims.(*JWTClaim)
	return claims, nil
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

func (s AuthenticationService) SignInOrSignUp(
	ctx context.Context, authenticationProvider string, identifier string, credential string, name string,
) (
	*ent.Authentication, error,
) {
	a, err := s.db.Authentication.Query().
		Where(
			authentication.ProviderEQ(authentication.Provider(authenticationProvider)),
			authentication.IdentifierEQ(identifier),
		).
		WithUser().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return s.SignUp(ctx, authenticationProvider, identifier, credential, name)
		}

		return nil, err
	}

	return a, nil
}

func (s AuthenticationService) SignUp(
	ctx context.Context, authenticationProvider string, identifier string, credential string, name string,
) (
	*ent.Authentication, error,
) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	userCreateQuery := s.db.User.Create()
	if name != "" {
		userCreateQuery = userCreateQuery.SetName(name)
	}

	u, err := userCreateQuery.Save(ctx)
	if err != nil {
		return nil, dbRollback(tx, err)
	}

	a, err := s.db.Authentication.Create().
		SetProvider(authentication.Provider(authenticationProvider)).
		SetIdentifier(identifier).
		SetCredential(credential).
		SetUserID(u.ID).
		Save(ctx)
	if err != nil {
		return nil, dbRollback(tx, err)
	}

	a.Edges.User = u

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return a, nil
}
