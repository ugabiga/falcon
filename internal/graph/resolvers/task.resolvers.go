package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input generated.CreateTaskInput) (*generated.Task, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	task, err := r.taskSrv.Create(ctx, claim.UserID, input)
	if err != nil {
		return nil, err
	}

	var respTask *generated.Task
	if err := deepcopy.CopyEx(&respTask, task); err != nil {
		return nil, err
	}

	return respTask, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, tradingAccountID string, taskID string, input generated.UpdateTaskInput) (*generated.Task, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	task, err := r.taskSrv.Update(ctx, claim.UserID, tradingAccountID, taskID, input)
	if err != nil {
		return nil, err
	}

	var respTask *generated.Task
	if err := deepcopy.CopyEx(&respTask, task); err != nil {
		return nil, err
	}

	return respTask, nil
}

func (r *queryResolver) TaskIndex(ctx context.Context, tradingAccountID *string) (*generated.TaskIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	tradingAccounts, err := r.tradingAccountSrv.GetByUserID(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return nil, err
	}

	var respTradingAccounts []*generated.TradingAccount
	for _, tradingAccount := range tradingAccounts {
		var ta *generated.TradingAccount
		if err := deepcopy.CopyEx(&ta, tradingAccount); err != nil {
			return nil, err
		}
		respTradingAccounts = append(respTradingAccounts, ta)
	}

	if len(respTradingAccounts) == 0 {
		return &generated.TaskIndex{}, nil
	}

	selectedRespTradingAccount := respTradingAccounts[0]

	if tradingAccountID != nil {
		for _, tradingAccount := range respTradingAccounts {
			if tradingAccount.ID == *tradingAccountID {
				selectedRespTradingAccount = tradingAccount
				break
			}
		}
	}

	tasks, err := r.taskSrv.GetByTradingAccount(ctx, selectedRespTradingAccount.ID)
	if err != nil {
		return nil, err
	}

	var respTasks []*generated.Task
	for _, task := range tasks {
		var t *generated.Task
		if err := deepcopy.CopyEx(&t, task); err != nil {
			return nil, err
		}
		respTasks = append(respTasks, t)
	}
	selectedRespTradingAccount.Tasks = respTasks

	return &generated.TaskIndex{
		TradingAccounts:        respTradingAccounts,
		SelectedTradingAccount: selectedRespTradingAccount,
	}, nil
}
