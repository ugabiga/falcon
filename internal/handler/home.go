package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
	"net/http"
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

func (h HomeHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "this is falcon",
	})
}
