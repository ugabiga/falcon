package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/helper"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/service"
	"net/http"
)

type TaskHandler struct {
	taskService           *service.TaskService
	tradingAccountService *service.TradingAccountService
}

func NewTaskHandler(
	taskService *service.TaskService,
	tradingAccountService *service.TradingAccountService,
) *TaskHandler {
	return &TaskHandler{
		taskService:           taskService,
		tradingAccountService: tradingAccountService,
	}
}

func (h TaskHandler) SetRoutes(e *echo.Group) {
	e.GET("/tasks", h.Index)
}

// Index godoc
//
//	@Summary		Get user tasks
//	@Description	Get user tasks
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			trading_account_id	query		string	false	"Trading account ID"
//	@Success		200					{object}	TaskIndexResponse
//	@Failure		400					{object}	handler.APIError
//	@Router			/api/v1/tasks [get]
func (h TaskHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	claim := helper.MustJWTClaim(c)
	tradingAccountID := c.QueryParam("trading_account_id")

	tradingAccounts, err := h.tradingAccountService.GetByUserID(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return err
	}

	if len(tradingAccounts) == 0 {
		return c.JSON(http.StatusOK, TaskIndexResponse{})
	}

	if tradingAccountID == "" {
		tradingAccountID = tradingAccounts[0].ID
	}

	tasks, err := h.taskService.GetByTradingAccount(ctx, tradingAccountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newTaskIndexResponse(tasks, tradingAccounts, tradingAccountID))
}

type TaskIndexResponse struct {
	TradingAccounts        []model.TradingAccount `json:"trading_accounts"`
	SelectedTradingAccount model.TradingAccount   `json:"selected_trading_account"`
	SelectedTasks          []model.Task           `json:"selected_tasks"`
}

func newTaskIndexResponse(tasks []model.Task, tradingAccounts []model.TradingAccount, tradingAccountID string) TaskIndexResponse {
	resp := TaskIndexResponse{}
	resp.SelectedTasks = tasks
	resp.TradingAccounts = tradingAccounts

	// Find selected trading account
	for _, tradingAccount := range tradingAccounts {
		if tradingAccount.ID == tradingAccountID {
			resp.SelectedTradingAccount = tradingAccount
			break
		}
	}

	return resp
}
