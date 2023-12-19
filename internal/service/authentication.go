package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/authentication"
)

type AuthenticationService struct {
	db *ent.Client
}

func NewAuthenticationService(db *ent.Client) *AuthenticationService {
	return &AuthenticationService{
		db: db,
	}
}

func (s AuthenticationService) SignInOrSignUp(
	ctx context.Context, authenticationProvider string, identifier string, credential string,
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
			return s.SignUp(ctx, authenticationProvider, identifier, credential)
		}

		return nil, err
	}

	return a, nil
}

func (s AuthenticationService) SignUp(
	ctx context.Context, authenticationProvider string, identifier string, credential string,
) (
	*ent.Authentication, error,
) {
	a, err := s.db.Authentication.Create().
		SetProvider(authentication.Provider(authenticationProvider)).
		SetIdentifier(identifier).
		SetCredential(credential).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	u, err := s.db.User.Create().
		AddAuthentications(a).
		Save(ctx)

	a.Edges.User = u

	if err != nil {
		return nil, err
	}

	return a, nil
}
