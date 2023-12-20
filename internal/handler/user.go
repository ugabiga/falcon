package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	userService *service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) SetRoutes(e *echo.Group) {
	e.GET("/user", h.Get)
}

func (h UserHandler) Get(c echo.Context) error {
	var userID uint64 = 1
	user, err := h.userService.GetByID(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
