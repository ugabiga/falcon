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

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) GenerateToken() (string, error) {
	claims := &jwtCustomClaims{
		"John Doe",
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

func (s *JWTService) Middleware(whiteList []string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey:  []byte(SecretKey),
		TokenLookup: "header:Authorization:Bearer ,cookie:falcon.access_token",
		Skipper: func(c echo.Context) bool {
			for _, v := range whiteList {
				if strings.HasPrefix(c.Request().RequestURI, v) {
					return true
				}
			}
			return false
		},
	})
}
func (s *JWTService) CustomClaimsName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims.Name
}
