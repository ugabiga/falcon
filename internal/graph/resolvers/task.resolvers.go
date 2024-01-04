package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/graph/generated"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input generated.CreateTaskInput) (*generated.Task, error) {
	//claim := helper.MustJWTClaimInResolver(ctx)

	//task, err := r.taskSrv.Create(ctx, claim.UserID, input)
	//if err != nil {
	//	return nil, err
	//}

	//return converter.ToTask(task)

	return nil, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input generated.UpdateTaskInput) (*generated.Task, error) {
	//r.logger.Printf("Reqest %v", debug.ToJSONStr(input))
	//
	//claim := helper.MustJWTClaimInResolver(ctx)
	//
	//task, err := r.taskSrv.Update(ctx, claim.UserID, id, input)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return converter.ToTask(task)
	return nil, nil
}

func (r *queryResolver) TaskIndex(ctx context.Context, tradingAccountID *string) (*generated.TaskIndex, error) {
	//claim := helper.MustJWTClaimInResolver(ctx)
	//
	//tradingAccounts, err := r.tradingAccountSrv.GetWithTask(
	//	ctx,
	//	claim.UserID,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//accounts, err := converter.ToTradingAccounts(tradingAccounts)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(accounts) == 0 {
	//	return &generated.TaskIndex{}, nil
	//}
	//
	//selectedAccount := accounts[0]
	//
	//if tradingAccountID != nil {
	//	for _, account := range accounts {
	//		if account.ID == *tradingAccountID {
	//			selectedAccount = account
	//			break
	//		}
	//	}
	//}
	//
	//return &generated.TaskIndex{
	//	TradingAccounts:        accounts,
	//	SelectedTradingAccount: selectedAccount,
	//}, nil
	//

	return nil, nil
}
