package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/ugabiga/falcon/internal/graph/converter"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/service"
)

func (r *queryResolver) User(ctx context.Context, id string, withOptions generated.UserWithOptions) (*generated.User, error) {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
	logger.Printf("User query with id: %s", id)
	logger.Printf("User query with options: %+v", withOptions)

	dummy, err := r.userSrv.CreateDummy(ctx,
		converter.StringToInt(id),
		service.UserWithOptions{
			WithAuthentications: withOptions.WithAuthentications,
			WithTradingAccounts: withOptions.WithTradingAccounts,
		},
	)
	if err != nil {
		return nil, err
	}

	user, err := converter.ToUser(dummy)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context, where generated.UserWhereInput) ([]*generated.User, error) {
	whereInput := converter.ToUserWhereInput(where)
	users, err := r.userSrv.Query(ctx, whereInput)
	if err != nil {
		return nil, err
	}

	return converter.ToUsers(users)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
