package handler

import (
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h HomeHandler) SetRoutes(e *echo.Group) {
	e.GET("/", h.Index)
}

type HomeIndex struct {
	Layout Layout
	Title  string
}

func (h HomeHandler) Index(c echo.Context) error {
	return RenderPage(
		c.Response().Writer,
		HomeIndex{
			Title: "Home Page",
		},
		"/index.html",
	)
}
