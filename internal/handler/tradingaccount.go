package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
)

type TradingAccountHandler struct {
}

func NewTradingAccountHandler() *TradingAccountHandler {
	return &TradingAccountHandler{}
}

func (h *TradingAccountHandler) SetRoutes(e *echo.Group) {
	e.GET("/tradingaccount", h.Index)
}

type TradingAccountIndex struct {
	Layout model.Layout
	Title  string
}

func (h *TradingAccountHandler) Index(c echo.Context) error {
	r := RenderPage(
		c.Response().Writer,
		TradingAccountIndex{
			Layout: middleware.ExtractLayout(c),
			Title:  "Task Page",
		},
		"/tradingaccount/index.html",
	)

	return r
}
