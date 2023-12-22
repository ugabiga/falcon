package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
)

type HomeHandler struct {
	authenticationService *service.AuthenticationService
}

func NewHomeHandler(
	authenticationService *service.AuthenticationService,
) *HomeHandler {
	return &HomeHandler{
		authenticationService: authenticationService,
	}
}

func (h HomeHandler) SetRoutes(e *echo.Group) {
	e.GET("/", h.Index)
}

type HomeIndex struct {
	Layout model.Layout
	Title  string
}

func (h HomeHandler) Index(c echo.Context) error {
	r := RenderPage(
		c.Response().Writer,
		HomeIndex{
			Layout: middleware.ExtractLayout(c),
			Title:  "Home Page",
		},
		"/index.html",
	)

	return r
}
