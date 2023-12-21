package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
	"net/http"
)

type AuthenticationHandler struct {
	jwtService *service.JWTService
}

func NewAuthenticationHandler(
	jwtService *service.JWTService,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		jwtService: jwtService,
	}
}

func (h AuthenticationHandler) SetRoutes(e *echo.Group) {
	e.GET("/auth/signin", h.SignIn)
	e.POST("/auth/action/_test", h.ActionTest)
}

type SignIn struct {
	Layout Layout
}

func (h AuthenticationHandler) SignIn(c echo.Context) error {
	return RenderPage(
		c.Response().Writer,
		SignIn{},
		"/auth/index.html",
	)
}

func (h AuthenticationHandler) ActionTest(c echo.Context) error {
	return RenderComponent(
		c.Response().Writer,
		SignIn{},
		"/refresh.html",
	)
}

func (h AuthenticationHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}

func (h AuthenticationHandler) Post(c echo.Context) error {
	t, err := h.jwtService.GenerateToken()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h AuthenticationHandler) Protected(c echo.Context) error {
	name := h.jwtService.CustomClaimsName(c)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + name + "!",
	})
}
