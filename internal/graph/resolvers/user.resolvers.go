package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/helper"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, input generated.UpdateUserInput) (*generated.User, error) {
	r.logger.Printf("Input: %+v", input)
	claim := helper.MustJWTClaimInResolver(ctx)

	updateUser, err := r.userSrv.Update(
		ctx,
		claim.UserID,
		&ent.User{
			Name:     input.Name,
			Timezone: input.Timezone,
		},
	)
	if err != nil {
		return nil, err
	}

	return converter.ToUser(updateUser)
}

func (r *queryResolver) User(ctx context.Context) (*generated.User, error) {
	claim := helper.MustJWTClaimInResolver(ctx)

	user, err := r.userSrv.GetByID(ctx, claim.UserID)
	if err != nil {
		return nil, err
	}

	return converter.ToUser(user)
}

func (r *queryResolver) Users(ctx context.Context, where generated.UserWhereInput) ([]*generated.User, error) {
	var whereInput ent.UserWhereInput
	converter.BindWhereInput(&where, &whereInput)

	users, err := r.userSrv.Query(ctx, whereInput)
	if err != nil {
		return nil, err
	}

	return converter.ToUsers(users)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
