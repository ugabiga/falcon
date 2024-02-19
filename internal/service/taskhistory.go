package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
)

type TaskHistoryService struct {
	repo *repository.DynamoRepository
}

func NewTaskHistoryService(
	tradingRepo *repository.DynamoRepository,
) *TaskHistoryService {
	return &TaskHistoryService{
		repo: tradingRepo,
	}
}

func (s *TaskHistoryService) GetTaskHistoryByTaskId(ctx context.Context, userID, tradingAccountID, taskID string) (*model.Task, []model.TaskHistory, error) {
	task, err := s.repo.GetTask(ctx, tradingAccountID, taskID)
	if err != nil {
		return nil, nil, err
	}

	if task.UserID != userID {
		return nil, nil, ErrUnAuthorizedAction
	}

	taskHistories, err := s.repo.GetTaskHistoriesByTaskID(ctx, task.ID)
	if err != nil {
		return nil, nil, err
	}

	return task, taskHistories, nil
}

func (s *TaskHistoryService) UpdateAllTaskHistoryTTL(ctx context.Context) error {
	histories, err := s.repo.GetAllTaskHistories(ctx)
	if err != nil {
		log.Printf("error getting all task histories: %v", err)
		return err
	}

	for _, th := range histories {
		_, err := s.repo.UpdateTaskHistory(ctx, *th)
		if err != nil {
			log.Printf("error updating task history: %v", err)
			return err
		}
	}

	return nil
}
