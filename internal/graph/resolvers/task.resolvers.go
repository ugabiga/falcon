package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *queryResolver) Tasks(ctx context.Context, tradingAccountID *string) ([]*generated.Task, error) {
	claim := helper.MustJWTClaimInResolver(ctx)
	if tradingAccountID == nil {
		tradingAccount, err := r.tradingAccountSrv.First(ctx, claim.UserID)
		if err != nil {
			return nil, err
		}

		tasks, err := r.taskSrv.GetByTradingAccount(ctx, tradingAccount.ID)
		if err != nil {
			return nil, err
		}

		return converter.ToTasks(tasks)
	}

	return nil, nil
}
