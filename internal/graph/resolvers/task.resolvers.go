package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/graph/model"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) CreateTask(ctx context.Context, tradingAccountID string, currency string, days string, hours string, typeArg string, params model.JSON) (*generated.Task, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	task, err := r.taskSrv.Create(
		ctx,
		claim.UserID,
		str.New(tradingAccountID).ToIntDefault(0),
		currency,
		days,
		hours,
		typeArg,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToTask(task)
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, currency string, days string, hours string, typeArg string, isActive bool, params model.JSON) (*generated.Task, error) {
	r.logger.Printf("Reqest %v", debug.ToJSONStr(map[string]interface{}{
		"id":       id,
		"currency": currency,
		"days":     days,
		"hours":    hours,
		"type":     typeArg,
		"isActive": isActive,
		"params":   params,
	}))

	claim := helper.MustJWTClaimInResolver(ctx)

	task, err := r.taskSrv.Update(
		ctx,
		claim.UserID,
		str.New(id).ToIntDefault(0),
		currency,
		days,
		hours,
		typeArg,
		isActive,
		params,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToTask(task)
}

func (r *queryResolver) TaskIndex(ctx context.Context, tradingAccountID *string) (*generated.TaskIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	tradingAccounts, err := r.tradingAccountSrv.GetWithTask(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return nil, err
	}

	accounts, err := converter.ToTradingAccounts(tradingAccounts)
	if err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		return &generated.TaskIndex{}, nil
	}

	selectedAccount := accounts[0]

	if tradingAccountID != nil {
		for _, account := range accounts {
			if account.ID == *tradingAccountID {
				selectedAccount = account
				break
			}
		}
	}

	return &generated.TaskIndex{
		TradingAccounts:        accounts,
		SelectedTradingAccount: selectedAccount,
	}, nil
}
