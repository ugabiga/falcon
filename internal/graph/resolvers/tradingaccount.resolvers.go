package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) CreateTradingAccount(ctx context.Context, name string, exchange string, key string, secret string) (*generated.TradingAccount, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	newTradingAccount, err := r.tradingAccountSrv.Create(
		ctx,
		claim.UserID,
		name,
		exchange,
		key,
		secret,
		"",
	)
	if err != nil {
		return nil, err
	}

	var resp *generated.TradingAccount
	if err := deepcopy.CopyEx(&resp, newTradingAccount); err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *mutationResolver) UpdateTradingAccount(ctx context.Context, id string, name *string, exchange *string, key *string, secret *string) (bool, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	err := r.tradingAccountSrv.Update(
		ctx,
		id,
		claim.UserID,
		name,
		exchange,
		key,
		secret,
		nil,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteTradingAccount(ctx context.Context, id string) (bool, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	err := r.tradingAccountSrv.Delete(
		ctx,
		claim.UserID,
		id,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) TradingAccountIndex(ctx context.Context) (*generated.TradingAccountIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	tradingAccounts, err := r.tradingAccountSrv.GetByUserID(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return nil, err
	}

	var resp []*generated.TradingAccount
	for _, tradingAccount := range tradingAccounts {
		var ta *generated.TradingAccount
		if err := deepcopy.CopyEx(&ta, tradingAccount); err != nil {
			return nil, err
		}
		resp = append(resp, ta)
	}

	//if len(resp) == 0 {
	//	return &generated.TradingAccountIndex{
	//		TradingAccounts: []*generated.TradingAccount{},
	//	}, nil
	//}

	return &generated.TradingAccountIndex{
		TradingAccounts: resp,
	}, nil
}
