package handler

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
)

const (
	LayoutContextKey = "layout"
)

type Layout struct {
	Claim   *service.JWTClaim
	IsLogin bool
}

type LayoutPage struct {
	Layout Layout
}

func LayoutMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (returnErr error) {
			name := "falcon.access_token"
			cookies := c.Cookies()
			if len(cookies) == 0 {
				return next(c)
			}

			tokenStr, err := c.Request().Cookie(name)
			if err != nil {
				return next(c)
			}

			token, err := jwt.ParseWithClaims(
				tokenStr.Value,
				&service.JWTClaim{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, err
					}
					return []byte(service.SecretKey), nil
				},
			)
			if err != nil {
				return next(c)
			}

			if claims, ok := token.Claims.(*service.JWTClaim); ok {
				c.Set(LayoutContextKey, Layout{
					Claim:   claims,
					IsLogin: true,
				})
			}
			return next(c)
		}
	}
}

func ExtractLayout(c echo.Context) Layout {
	layout, ok := c.Get(LayoutContextKey).(Layout)
	if !ok {
		return Layout{}
	}

	return layout
}
