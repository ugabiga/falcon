package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type TaskHistoryService struct {
	db              *ent.Client
	taskHistoryRepo *repository.TaskHistoryDynamoRepository
}

func NewTaskHistoryService(db *ent.Client, taskHistoryRepo *repository.TaskHistoryDynamoRepository) *TaskHistoryService {
	return &TaskHistoryService{
		db:              db,
		taskHistoryRepo: taskHistoryRepo,
	}
}

func (s *TaskHistoryService) GetTaskHistoryByTaskId(ctx context.Context, taskId string) ([]model.TaskHistory, error) {
	return s.taskHistoryRepo.GetByTaskID(
		ctx,
		taskId,
	)
}
