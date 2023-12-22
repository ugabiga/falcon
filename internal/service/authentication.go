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
	SecretKey   = "secret"
	maxAge      = 86400 * 30
	baseURL     = "http://localhost:3000"
	callbackURL = baseURL + "/auth/signin/google/callback"
)

type JWTClaim struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
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
	key := s.cfg.SessionSecretKey
	googleClientID := s.cfg.GoogleClientID
	googleClientSecret := s.cfg.GoogleClientSecret

	store := sessions.NewCookieStore([]byte(key))
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

func (s AuthenticationService) JWTToken(userID int, name string, isAdmin bool) (string, error) {
	claims := &JWTClaim{
		userID,
		name,
		isAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s AuthenticationService) JWTClaim(c echo.Context) *JWTClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTClaim)
	return claims
}

func (s AuthenticationService) JWTMiddleware(matchWhiteList, prefixWhiteList []string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaim)
		},
		SigningKey:  []byte(SecretKey),
		TokenLookup: "header:Authorization:Bearer ,cookie:falcon.access_token",
		Skipper: func(c echo.Context) bool {
			for _, v := range matchWhiteList {
				if c.Path() == v {
					return true
				}
			}

			for _, v := range prefixWhiteList {
				if strings.HasPrefix(c.Path(), v) {
					return true
				}
			}
			return false
		},
	})
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
