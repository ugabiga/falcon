package handler

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
	"log"
	"net/http"
	"time"
)

type AuthenticationHandler struct {
	authenticationService *service.AuthenticationService
}

func NewAuthenticationHandler(
	authenticationService *service.AuthenticationService,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: authenticationService,
	}
}

func (h AuthenticationHandler) SetRoutes(e *echo.Group) {
	e.GET("/auth/signin", h.SignInIndex)
	e.GET("/auth/signin/:provider", h.SignIn)
	e.GET("/auth/signin/:provider/callback", h.SignInCallback)
	e.GET("/auth/signout/:provider", h.SignOut)
	e.GET("/auth/protected", h.Protected)
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

	a, err := h.authenticationService.SignInOrSignUp(
		c.Request().Context(),
		"google",
		user.UserID,
		user.AccessToken,
		user.Name,
	)
	if err != nil {
		return err
	}

	token, err := h.authenticationService.JWTToken(
		a.UserID,
		a.Edges.User.Name,
		false,
	)
	if err != nil {
		return err
	}

	log.Printf("User: %+v", debug.ToJSONStr(user))

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     "falcon.access_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(&cookie)

	return c.Redirect(http.StatusFound, "/")
}

func (h AuthenticationHandler) SignOut(c echo.Context) error {
	c.SetRequest(c.Request().WithContext(
		context.WithValue(
			c.Request().Context(),
			"provider",
			c.Param("provider"),
		),
	))
	if err := gothic.Logout(c.Response(), c.Request()); err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

type SignInIndex struct {
	Layout model.Layout
}

func (h AuthenticationHandler) SignInIndex(c echo.Context) error {
	return RenderPage(
		c.Response().Writer,
		SignInIndex{},
		"/auth/index.html",
	)
}

func (h AuthenticationHandler) ActionTest(c echo.Context) error {
	return RenderComponent(
		c.Response().Writer,
		SignInIndex{},
		"/refresh.html",
	)
}

func (h AuthenticationHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}

func (h AuthenticationHandler) Protected(c echo.Context) error {
	claim := h.authenticationService.JWTClaim(c)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + claim.Name + "!",
	})
}
