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
	taskHistories, err := r.taskHistorySrv.GetTaskHistoryByTaskId(
		ctx,
		str.New(taskID).ToIntDefault(0),
	)
	if err != nil {
		return nil, err
	}

	convertedTaskHistories, err := converter.ToTaskHistories(taskHistories)
	if err != nil {
		return nil, err
	}

	return &generated.TaskHistoryIndex{
		TaskHistories: convertedTaskHistories,
	}, nil
}
