package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/service"
	"log"
	"net/http"
)

type AuthenticationHandler struct {
	jwtService            *service.JWTService
	authenticationService *service.AuthenticationService
}

func NewAuthenticationHandler(
	jwtService *service.JWTService,
	authenticationService *service.AuthenticationService,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		jwtService:            jwtService,
		authenticationService: authenticationService,
	}
}

func (h AuthenticationHandler) SetRoutes(e *echo.Group) {
	e.GET("/auth/signin", h.SignInIndex)
	e.POST("/auth/action/_test", h.ActionTest)
	e.GET("/auth/signin/:provider", h.SignIn)
	e.GET("/auth/signin/:provider/callback", h.SignInCallback)
}

type SignIn struct {
	Layout Layout
}

func (h AuthenticationHandler) SignIn(c echo.Context) error {
	newRequestContext := c.Request().WithContext(
		context.WithValue(
			c.Request().Context(),
			"provider",
			c.Param("provider"),
		),
	)
	c.SetRequest(newRequestContext)

	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		log.Printf("User: %+v", gothUser)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		gothic.BeginAuthHandler(c.Response(), c.Request())
	}

	return nil
}

func (h AuthenticationHandler) SignInCallback(c echo.Context) error {
	log.Printf("in ProviderCallBack")
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		log.Printf("err: %+v", err)
		return err
	}

	log.Printf("User: %+v", debug.ToJSONStr(user))

	return c.Redirect(http.StatusFound, "/")
}

func (h AuthenticationHandler) SignInIndex(c echo.Context) error {
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
