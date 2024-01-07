package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type TaskHistoryService struct {
	db          *ent.Client
	tradingRepo *repository.DynamoRepository
}

func NewTaskHistoryService(
	db *ent.Client,
	tradingRepo *repository.DynamoRepository,
) *TaskHistoryService {
	return &TaskHistoryService{
		db:          db,
		tradingRepo: tradingRepo,
	}
}

func (s *TaskHistoryService) GetTaskHistoryByTaskId(ctx context.Context, userID, tradingAccountID, taskID string) (*model.Task, []model.TaskHistory, error) {
	task, err := s.tradingRepo.GetTask(ctx, tradingAccountID, taskID)
	if err != nil {
		return nil, nil, err
	}

	if task.UserID != userID {
		return nil, nil, ErrDoNotHaveAccess
	}

	taskHistories, err := s.tradingRepo.GetTaskHistoriesByTaskID(ctx, task.ID)
	if err != nil {
		return nil, nil, err
	}

	return task, taskHistories, nil
}
