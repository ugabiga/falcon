package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/handler/v1/request"
	"github.com/ugabiga/falcon/internal/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) SetRoutes(e *echo.Group) {
	e.GET("/users/me", h.Me)
	e.PUT("/users/me", h.Update)
}

// Me godoc
//
//	@Summary		Get user profile
//	@Description	Get user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	model.User
//	@Failure		401	{object}	handler.APIError
//	@Router			/api/v1/users/me [get]
func (h UserHandler) Me(c echo.Context) error {
	claim := helper.MustJWTClaim(c)

	user, err := h.userService.GetByID(c.Request().Context(), claim.UserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// Update godoc
//
//	@Summary		Update user profile
//	@Description	Update user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			req	body		request.UpdateUserRequest	true	"User profile"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	handler.APIError
//	@Failure		401	{object}	handler.APIError
//	@Router			/api/v1/users/me [put]
func (h UserHandler) Update(c echo.Context) error {
	claim := helper.MustJWTClaim(c)
	req := new(request.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := h.userService.UpdateV1(
		c.Request().Context(),
		claim.UserID,
		*req,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
