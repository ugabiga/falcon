package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
	"github.com/ugabiga/falcon/internal/ent/tradingaccount"
)

type TaskService struct {
	db *ent.Client
}

func NewTaskService(db *ent.Client) *TaskService {
	return &TaskService{db: db}
}

func (s TaskService) GetByTradingAccount(ctx context.Context, tradingAccountID int) ([]*ent.Task, error) {
	return s.db.Task.Query().
		Where(
			task.TradingAccountID(tradingAccountID),
		).
		WithTradingAccount().
		All(ctx)
}

func (s TaskService) Create(ctx context.Context, userID int, tradingAccountID int, cron string, typeArg string) (*ent.Task, error) {
	tradingAccount, err := s.db.TradingAccount.Query().
		Where(
			tradingaccount.UserID(userID),
			tradingaccount.ID(tradingAccountID),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return s.db.Task.Create().
		SetCron(cron).
		SetType(typeArg).
		SetTradingAccountID(tradingAccount.ID).
		Save(ctx)
}
