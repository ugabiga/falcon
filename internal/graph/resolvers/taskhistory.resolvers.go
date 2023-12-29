package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *queryResolver) TaskHistoryIndex(ctx context.Context, taskID int) (*generated.TaskHistoryIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	task, err := r.taskSrv.GetWithTaskHistory(
		ctx,
		claim.UserID,
		taskID,
	)
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
