package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
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
