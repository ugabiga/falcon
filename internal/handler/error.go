package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIError struct {
	Error   string
	Message string
}

type Error struct {
	Code    int
	Message string
}

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h ErrorHandler) ErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	ok := errors.As(err, &he)
	if ok {
		if he.Internal != nil {
			var herr *echo.HTTPError
			if errors.As(he.Internal, &herr) {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = APIError{
			Error:   http.StatusText(he.Code),
			Message: he.Message.(string),
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}

func (h ErrorHandler) DebugErrorHandler(err error, c echo.Context) {
	code := http.StatusBadRequest
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	c.Logger().Error(err)
	c.Response().WriteHeader(code)
	if err := c.JSON(code, Error{
		Code:    code,
		Message: err.Error(),
	}); err != nil {
		return
	}

	//if err := RenderPage(
	//	c.Response().Writer,
	//	ErrorIndex{
	//		Layout: middleware.ExtractLayout(c),
	//		Error: Error{
	//			Code:    code,
	//			Message: err.Error(),
	//		},
	//	},
	//	"/error/index.html",
	//); err != nil {
	//	c.Logger().Error(err)
	//}
}
