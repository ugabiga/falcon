package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) CreateTradingAccount(ctx context.Context, exchange string, currency string, identifier string, credential string) (*generated.TradingAccount, error) {
	r.logger.Printf("CreateTradingAccount: exchange=%s, currency=%s, identifier=%s", exchange, currency, identifier)

	claim := helper.MustJWTClaimInResolver(ctx)
	newTradingAccount, err := r.tradingAccountSrv.Create(
		ctx,
		claim.UserID,
		exchange,
		currency,
		identifier,
		credential,
		"",
	)
	if err != nil {
		return nil, err
	}

	return converter.ToTradingAccount(newTradingAccount)
}

func (r *mutationResolver) UpdateTradingAccount(ctx context.Context, id string, exchange *string, currency *string, identifier *string, credential *string) (bool, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	err := r.tradingAccountSrv.Update(
		ctx,
		str.New(id).ToIntDefault(0),
		claim.UserID,
		exchange,
		currency,
		identifier,
		credential,
		nil,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteTradingAccount(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TradingAccounts(ctx context.Context) ([]*generated.TradingAccount, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	tradingAccounts, err := r.tradingAccountSrv.Get(
		ctx,
		claim.UserID,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToTradingAccounts(tradingAccounts)
}

func (r *queryResolver) TradingAccountsWithTasks(ctx context.Context, id *string) (*generated.TradingAccountWithTasks, error) {
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
		return &generated.TradingAccountWithTasks{}, nil
	}

	selectedAccount := accounts[0]

	if id != nil {
		for _, account := range accounts {
			if account.ID == *id {
				selectedAccount = account
				break
			}
		}
	}

	return &generated.TradingAccountWithTasks{
		TradingAccounts:        accounts,
		SelectedTradingAccount: selectedAccount,
	}, nil
}
