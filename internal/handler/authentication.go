package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/service"
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
	e.POST("/auth/signin", h.Signin)
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

// Signin godoc
//
//	@Summary		Sign in
//	@Description	Sign in with OAuth
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			req	body		SignInRequest	true	"Sign in request"
//	@Success		200	{object}	SignInResponse
//	@Failure		500	{object}	handler.APIError
//	@Router			/auth/signin [post]
func (h AuthenticationHandler) Signin(c echo.Context) error {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	providerUser, err := h.authenticationService.VerifyOauthUser(
		c.Request().Context(),
		req.Type,
		req.AccessToken,
	)
	if err != nil {
		return err
	}

	a, u, err := h.authenticationService.SignInOrSignUp(
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
		u.Name,
		false,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, SignInResponse{
		Token: token,
	})
}

type ProtectedResponse struct {
	Message string `json:"message"`
}

// Protected godoc
//
//	@Summary		Protected
//	@Description	Protected route
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	ProtectedResponse
//	@Failure		401	{object}	handler.APIError
//	@Router			/auth/protected [get]
func (h AuthenticationHandler) Protected(c echo.Context) error {
	claim := helper.MustJWTClaim(c)
	return c.JSON(http.StatusOK, ProtectedResponse{
		Message: "Hello, " + claim.Name,
	})
}
