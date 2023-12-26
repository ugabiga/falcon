package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"time"
)

var (
	defaultTimeout time.Duration = 10 * time.Second
)

// Email is a struct that provides information on whether an email is verified or is the primary email address
type Email struct {
	Email    string
	Verified bool
	Primary  bool
}

type Claims struct {
	// Reserved claims
	Issuer  string  `json:"iss,omitempty" structs:"iss,omitempty"`
	Subject string  `json:"sub,omitempty" structs:"sub,omitempty"`
	Aud     string  `json:"aud,omitempty" structs:"aud,omitempty"`
	Iat     float64 `json:"iat,omitempty" structs:"iat,omitempty"`
	Exp     float64 `json:"exp,omitempty" structs:"exp,omitempty"`

	// Default profile claims
	Name              string `json:"name,omitempty" structs:"name,omitempty"`
	FamilyName        string `json:"family_name,omitempty" structs:"family_name,omitempty"`
	GivenName         string `json:"given_name,omitempty" structs:"given_name,omitempty"`
	MiddleName        string `json:"middle_name,omitempty" structs:"middle_name,omitempty"`
	NickName          string `json:"nickname,omitempty" structs:"nickname,omitempty"`
	PreferredUsername string `json:"preferred_username,omitempty" structs:"preferred_username,omitempty"`
	Profile           string `json:"profile,omitempty" structs:"profile,omitempty"`
	Picture           string `json:"picture,omitempty" structs:"picture,omitempty"`
	Website           string `json:"website,omitempty" structs:"website,omitempty"`
	Gender            string `json:"gender,omitempty" structs:"gender,omitempty"`
	Birthdate         string `json:"birthdate,omitempty" structs:"birthdate,omitempty"`
	ZoneInfo          string `json:"zoneinfo,omitempty" structs:"zoneinfo,omitempty"`
	Locale            string `json:"locale,omitempty" structs:"locale,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty" structs:"updated_at,omitempty"`
	Email             string `json:"email,omitempty" structs:"email,omitempty"`
	EmailVerified     bool   `json:"email_verified,omitempty" structs:"email_verified"`
	Phone             string `json:"phone,omitempty" structs:"phone,omitempty"`
	PhoneVerified     bool   `json:"phone_verified,omitempty" structs:"phone_verified"`

	// Custom profile claims that are provider specific
	CustomClaims map[string]interface{} `json:"custom_claims,omitempty" structs:"custom_claims,omitempty"`

	// TODO: Deprecate in next major release
	FullName    string `json:"full_name,omitempty" structs:"full_name,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty" structs:"avatar_url,omitempty"`
	Slug        string `json:"slug,omitempty" structs:"slug,omitempty"`
	ProviderId  string `json:"provider_id,omitempty" structs:"provider_id,omitempty"`
	UserNameKey string `json:"user_name,omitempty" structs:"user_name,omitempty"`
}

// UserProvidedData is a struct that contains the user's data returned from the oauth provider
type UserProvidedData struct {
	Emails   []Email
	Metadata *Claims
}

type OAuthProvider interface {
	GetUserData(context.Context, *oauth2.Token) (*UserProvidedData, error)
}

func SafeClose(closer io.Closer) {
	if err := closer.Close(); err != nil {
		logrus.WithError(err).Warn("Close operation failed")
	}
}

func makeRequest(ctx context.Context, tok *oauth2.Token, g *oauth2.Config, url string, dst interface{}) error {
	client := g.Client(ctx, tok)
	client.Timeout = defaultTimeout
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer SafeClose(res.Body)

	bodyBytes, _ := io.ReadAll(res.Body)
	defer SafeClose(res.Body)
	res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}
