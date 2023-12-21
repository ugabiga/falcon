package service

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

const (
	SecretKey = "secret"
)

type JWTClaim struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}
type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) GenerateToken(userID int, name string, isAdmin bool) (string, error) {
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

func (s *JWTService) GenerateDummyToken() (string, error) {
	claims := &JWTClaim{
		1,
		"dummy",
		true,
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

func (s *JWTService) Middleware(matchWhiteList, prefixWhiteList []string) echo.MiddlewareFunc {
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

func (s *JWTService) Claim(c echo.Context) *JWTClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTClaim)
	return claims
}
