package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
	"log"
	"net/http"
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
	e.POST("/auth/signin", h.NewSignIn)
	e.GET("/auth/protected", h.Protected)
	e.GET("/auth/signin", h.SignInIndex)
	e.GET("/auth/signin/:provider", h.SignIn)
	e.GET("/auth/signin/:provider/callback", h.SignInCallback)
	e.GET("/auth/signout/:provider", h.SignOut)
}

type SignInRequest struct {
	Type        string `json:"type"`
	AccountID   string `json:"account_id"`
	AccessToken string `json:"access_token"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func (h AuthenticationHandler) NewSignIn(c echo.Context) error {
	log.Printf("Request: %+v", c.Request())
	//request get token in json

	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	log.Printf("Request: %+v", req)

	providerUser, err := h.authenticationService.VerifyUser(
		c.Request().Context(),
		req.Type,
		req.AccessToken,
	)
	if err != nil {
		return err
	}

	log.Printf("User: %+v", debug.ToJSONStr(providerUser))

	a, err := h.authenticationService.SignInOrSignUp(
		c.Request().Context(),
		req.Type,
		req.AccountID,
		req.AccessToken,
		providerUser.Metadata.Name,
	)
	if err != nil {
		return err
	}

	token, err := h.authenticationService.CreateJWTToken(
		a.UserID,
		a.Edges.User.Name,
		false,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, SignInResponse{
		Token: token,
	})
}

func (h AuthenticationHandler) Protected(c echo.Context) error {
	claim := helper.MustJWTClaim(c)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + claim.Name + "!",
	})
}

type SignInIndex struct {
	Layout model.Layout
}

func (h AuthenticationHandler) SignInIndex(c echo.Context) error {
	return RenderPage(
		c.Response().Writer,
		SignInIndex{
			Layout: middleware.ExtractLayout(c),
		},
		"/auth/index.html",
	)
}

func (h AuthenticationHandler) SignIn(c echo.Context) error {
	c.SetRequest(c.Request().WithContext(
		context.WithValue(
			c.Request().Context(),
			"provider",
			c.Param("provider"),
		),
	))

	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		log.Printf("User: %+v", gothUser)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		gothic.BeginAuthHandler(c.Response(), c.Request())
	}

	return nil
}

func (h AuthenticationHandler) SignInCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
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

	token, err := h.authenticationService.CreateJWTToken(
		a.UserID,
		a.Edges.User.Name,
		false,
	)
	if err != nil {
		return err
	}

	if err := helper.SetSession(c); err != nil {
		return err
	}

	helper.SetCookie(c, service.JWTCookieName, token)

	return c.Redirect(http.StatusFound, "http://localhost:3000")
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

	helper.RemoveCookie(c, service.JWTCookieName)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
