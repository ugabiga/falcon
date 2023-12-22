package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"net/http"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

type Error struct {
	Code    int
	Message string
}
type ErrorIndex struct {
	Layout model.Layout
	Error  Error
}

func (h ErrorHandler) DebugErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	c.Logger().Error(err)

	c.Response().WriteHeader(code)
	if err := RenderPage(
		c.Response().Writer,
		ErrorIndex{
			Layout: middleware.ExtractLayout(c),
			Error: Error{
				Code:    code,
				Message: err.Error(),
			},
		},
		"/error/index.html",
	); err != nil {
		c.Logger().Error(err)
	}
}
