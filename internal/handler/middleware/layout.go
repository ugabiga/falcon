package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
)

const (
	LayoutContextKey = "layout"
)

func LayoutMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (returnErr error) {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return next(c)
			}

			if claims, ok := token.Claims.(*service.JWTClaim); ok {
				c.Set(LayoutContextKey, model.Layout{
					Claim:   claims,
					IsLogin: true,
				})
			}
			return next(c)
		}
	}
}

func ExtractLayout(c echo.Context) model.Layout {
	layout, ok := c.Get(LayoutContextKey).(model.Layout)
	if !ok {
		return model.Layout{}
	}

	return layout
}
