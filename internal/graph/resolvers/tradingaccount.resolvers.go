package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) CreateTradingAccount(ctx context.Context, name string, exchange string, identifier string, credential string) (*generated.TradingAccount, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	newTradingAccount, err := r.tradingAccountSrv.Create(
		ctx,
		claim.UserID,
		name,
		exchange,
		identifier,
		credential,
		"",
	)
	if err != nil {
		return nil, err
	}

	return converter.ToTradingAccount(newTradingAccount)
}

func (r *mutationResolver) UpdateTradingAccount(ctx context.Context, id string, name *string, exchange *string, identifier *string, credential *string) (bool, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	err := r.tradingAccountSrv.Update(
		ctx,
		str.New(id).ToIntDefault(0),
		claim.UserID,
		name,
		exchange,
		identifier,
		credential,
		nil,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) TradingAccountIndex(ctx context.Context) (*generated.TradingAccountIndex, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	tradingAccounts, err := r.tradingAccountSrv.Get(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return nil, err
	}

	respTradingAccounts, err := converter.ToTradingAccounts(tradingAccounts)
	if err != nil {
		return nil, err
	}

	return &generated.TradingAccountIndex{
		TradingAccounts: respTradingAccounts,
	}, nil
}
