package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
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
	e.GET("/user", h.Index)
	e.POST("/user", h.Edit)
}

type UserIndex struct {
	Layout model.Layout
	User   *ent.User
}

func (h UserHandler) Index(c echo.Context) error {
	claim := helper.MustJWTClaim(c)

	user, err := h.userService.GetByID(
		c.Request().Context(),
		claim.UserID,
	)
	if err != nil {
		return err
	}

	return RenderPage(
		c.Response().Writer,
		UserIndex{
			Layout: middleware.ExtractLayout(c),
			User:   user,
		},
		"/user/index.html",
	)
}

type UserEditForm struct {
	Name     string `form:"name"`
	Timezone string `form:"timezone"`
}

func (h UserHandler) Edit(c echo.Context) error {
	claim := helper.MustJWTClaim(c)

	var form UserEditForm
	if err := c.Bind(&form); err != nil {
		return err
	}

	_, err := h.userService.Update(
		c.Request().Context(),
		claim.UserID,
		&ent.User{
			Name:     form.Name,
			Timezone: form.Timezone,
		},
	)
	if err != nil {
		return err
	}

	helper.NewToastEvent(c, "success", "Successfully update user")

	return nil
}
