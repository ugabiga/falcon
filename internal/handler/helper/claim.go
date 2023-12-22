package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
)

func JWTClaim(c echo.Context) (*service.JWTClaim, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, echo.ErrUnauthorized
	}
	claims := user.Claims.(*service.JWTClaim)
	return claims, nil
}

func MustJWTClaim(c echo.Context) *service.JWTClaim {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

	claims := user.Claims.(*service.JWTClaim)

	return claims
}
