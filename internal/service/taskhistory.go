package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/taskhistory"
)

type TaskHistoryService struct {
	db *ent.Client
}

func NewTaskHistoryService(db *ent.Client) *TaskHistoryService {
	return &TaskHistoryService{db: db}
}

func (s *TaskHistoryService) GetTaskHistoryByTaskId(ctx context.Context, taskId int) ([]*ent.TaskHistory, error) {
	return s.db.TaskHistory.Query().
		Where(taskhistory.TaskID(taskId)).
		WithTask().
		All(ctx)
}
