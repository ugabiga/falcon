package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
)

type HomeHandler struct {
	jwtService            *service.JWTService
	authenticationService *service.AuthenticationService
}

func NewHomeHandler(
	jwtService *service.JWTService,
	authenticationService *service.AuthenticationService,
) *HomeHandler {
	return &HomeHandler{
		jwtService:            jwtService,
		authenticationService: authenticationService,
	}
}

func (h HomeHandler) SetRoutes(e *echo.Group) {
	e.GET("/", h.Index)
}

type HomeIndex struct {
	Layout Layout
	Title  string
}

func (h HomeHandler) Index(c echo.Context) error {
	r := RenderPage(
		c.Response().Writer,
		HomeIndex{
			Layout: ExtractLayout(c),
			Title:  "Home Page",
		},
		"/index.html",
	)

	return r
}
