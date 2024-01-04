package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/handler/helper"
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
