package helper

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
)

func MustJWTClaim(c echo.Context) *service.JWTClaim {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

	claims := user.Claims.(*service.JWTClaim)

	return claims
}
func MustJWTClaimInResolver(c context.Context) *service.JWTClaim {
	user, ok := c.Value("user").(*jwt.Token)
	if !ok {
		return nil
	}

	claims := user.Claims.(*service.JWTClaim)

	return claims
}
func NewJWTClaimContext(c echo.Context) context.Context {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

	ctx := context.WithValue(c.Request().Context(), "user", user)

	return ctx
}
