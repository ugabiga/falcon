package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/graph/generated"
)

func (r *queryResolver) TaskHistoryIndex(ctx context.Context, taskID string) (*generated.TaskHistoryIndex, error) {
	//claim := helper.MustJWTClaimInResolver(ctx)
	//
	//task, err := r.taskSrv.GetWithTaskHistory(
	//	ctx,
	//	claim.UserID,
	//	taskID,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//convertedTask, err := converter.ToTask(task)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &generated.TaskHistoryIndex{
	//	Task:          convertedTask,
	//	TaskHistories: convertedTask.TaskHistories,
	//}, nil

	return nil, nil
}
