package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/handler/helper"

	"github.com/ugabiga/falcon/internal/graph/generated"
)

func (r *queryResolver) TaskHistoryIndex(ctx context.Context, taskID string) (*generated.TaskHistoryIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	task, err := r.taskSrv.Get(ctx, claim.UserID, taskID)
	if err != nil {
		return nil, err
	}

	var respTask *generated.Task
	if err := deepcopy.CopyEx(&respTask, task); err != nil {
		return nil, err
	}

	taskHistories, err := r.taskHistorySrv.GetTaskHistoryByTaskId(ctx, taskID)
	if err != nil {
		return nil, err
	}

	var respTaskHistories []*generated.TaskHistory
	for _, taskHistory := range taskHistories {
		var th *generated.TaskHistory
		if err := deepcopy.CopyEx(&th, taskHistory); err != nil {
			return nil, err
		}
		respTaskHistories = append(respTaskHistories, th)
	}

	return &generated.TaskHistoryIndex{
		Task:          respTask,
		TaskHistories: respTaskHistories,
	}, nil
}
