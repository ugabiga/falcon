package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
)

type TradingAccountHandler struct {
	tradingAccountService *service.TradingAccountService
}

func NewTradingAccountHandler(
	tradingAccountService *service.TradingAccountService,
) *TradingAccountHandler {
	return &TradingAccountHandler{
		tradingAccountService: tradingAccountService,
	}
}

func (h *TradingAccountHandler) SetRoutes(e *echo.Group) {
	e.GET("/tradingaccount", h.Index)
	e.POST("/tradingaccount", h.Add)
	e.GET("/tradingaccount/list", h.List)
}

type TradingAccountIndex struct {
	Layout          model.Layout
	Title           string
	TradingAccounts []*ent.TradingAccount
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

type TradingAccountList struct {
	TradingAccounts []*ent.TradingAccount
}

func (h *TradingAccountHandler) List(c echo.Context) error {
	claim := helper.MustJWTClaim(c)
	tradingAccounts, err := h.tradingAccountService.Get(
		c.Request().Context(),
		claim.UserID,
	)
	if err != nil {
		return err
	}

	for _, tradingAccount := range tradingAccounts {
		tradingAccount.Exchange = str.ToCamel(tradingAccount.Exchange)
	}

	return RenderComponent(
		c.Response().Writer,
		TradingAccountList{
			TradingAccounts: tradingAccounts,
		},
		"/tradingaccount/list.html",
	)
}

type TradingAccountAddForm struct {
	Exchange   string `form:"exchange"`
	Currency   string `form:"currency"`
	Identifier string `form:"identifier"`
	Credential string `form:"credential"`
}

type TradingAccountAdd struct {
	TradingAccount *ent.TradingAccount
}

func (h *TradingAccountHandler) Add(c echo.Context) error {
	var form TradingAccountAddForm
	if err := c.Bind(&form); err != nil {
		return err
	}

	claim := helper.MustJWTClaim(c)
	tradingAccount, err := h.tradingAccountService.Create(
		c.Request().Context(),
		claim.UserID,
		form.Exchange,
		form.Exchange,
		form.Identifier,
		form.Credential,
		"",
	)
	if err != nil {
		return err
	}

	return RenderComponent(
		c.Response().Writer,
		TradingAccountAdd{
			TradingAccount: tradingAccount,
		},
		"/tradingaccount/add.html",
	)
}
