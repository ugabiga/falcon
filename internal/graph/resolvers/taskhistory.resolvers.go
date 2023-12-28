package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func (r *queryResolver) TaskHistoryIndex(ctx context.Context, taskID string) (*generated.TaskHistoryIndex, error) {
	task, err := r.taskSrv.GetWithTaskHistory(ctx, str.New(taskID).ToIntDefault(0))
	if err != nil {
		return nil, err
	}

	convertedTask, err := converter.ToTask(task)
	if err != nil {
		return nil, err
	}

	return &generated.TaskHistoryIndex{
		Task:          convertedTask,
		TaskHistories: convertedTask.TaskHistories,
	}, nil
}
