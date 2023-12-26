package authentication

import (
	"context"
	"golang.org/x/oauth2"
)

const (
	UserInfoEndpointGoogle      = "https://www.googleapis.com/userinfo/v2/me"
	AuthorizationEndpointGoogle = "https://accounts.google.com/o/oauth2/auth"
)

var (
	internalUserInfoEndpointGoogle = UserInfoEndpointGoogle
)

type googleUser struct {
	ID            string `json:"id"`
	Subject       string `json:"sub"`
	Issuer        string `json:"iss"`
	Name          string `json:"name"`
	AvatarURL     string `json:"picture"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	EmailVerified bool   `json:"email_verified"`
	HostedDomain  string `json:"hd"`
}

func (u googleUser) IsEmailVerified() bool {
	return u.VerifiedEmail || u.EmailVerified
}

type googleProvider struct {
	config *oauth2.Config
}

func NewGoogleProvider(googleClientID, googleClientSecret string) (OAuthProvider, error) {
	return &googleProvider{
		config: &oauth2.Config{
			ClientID:     googleClientID,
			ClientSecret: googleClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL: AuthorizationEndpointGoogle,
			},
		},
	}, nil
}

func (g googleProvider) GetUserData(ctx context.Context, token *oauth2.Token) (*UserProvidedData, error) {
	var u googleUser
	if err := makeRequest(ctx, token, g.config, internalUserInfoEndpointGoogle, &u); err != nil {
		return nil, err
	}

	return &UserProvidedData{
		Emails: []Email{
			{
				Email:    u.Email,
				Verified: u.IsEmailVerified(),
				Primary:  true,
			},
		},
		Metadata: &Claims{
			Subject:   u.Subject,
			Issuer:    u.Issuer,
			Name:      u.Name,
			AvatarURL: u.AvatarURL,
			Email:     u.Email,
		},
	}, nil
}
