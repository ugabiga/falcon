package service

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/authentication"
)

type AuthenticationService struct {
	db  *ent.Client
	cfg *config.Config
}

func NewAuthenticationService(db *ent.Client, cfg *config.Config) *AuthenticationService {
	a := &AuthenticationService{
		db:  db,
		cfg: cfg,
	}
	a.InitializeProviders()

	return a
}

func (s AuthenticationService) InitializeProviders() {
	MaxAge := 86400 * 30
	baseURL := "http://localhost:3000"
	callbackURL := baseURL + "/auth/signin/google/callback"
	key := s.cfg.SessionSecretKey
	googleClientID := s.cfg.GoogleClientID
	googleClientSecret := s.cfg.GoogleClientSecret

	store := sessions.NewCookieStore([]byte(key))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   MaxAge,
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
