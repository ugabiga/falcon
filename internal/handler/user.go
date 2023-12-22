package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
	"log"
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
}

type UserIndex struct {
	Layout model.Layout
	User   *ent.User
}

func (h UserHandler) Index(c echo.Context) error {
	claim := helper.MustJWTClaim(c)
	log.Printf("claim: %+v", debug.ToJSONStr(claim))
	user, err := h.userService.GetByID(
		c.Request().Context(),
		claim.UserID,
	)
	if err != nil {
		return err
	}
	log.Printf("user: %+v", debug.ToJSONStr(user))

	return RenderPage(
		c.Response().Writer,
		UserIndex{
			Layout: middleware.ExtractLayout(c),
			User:   user,
		},
		"/user/index.html",
	)
}
