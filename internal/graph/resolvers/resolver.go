package resolvers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rs/zerolog"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/service"
	"os"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userSrv           *service.UserService
	tradingAccountSrv *service.TradingAccountService
	taskSrv           *service.TaskService
	logger            zerolog.Logger
}

func NewResolver(
	userSrv *service.UserService,
	tradingAccountSrv *service.TradingAccountService,
	taskSrv *service.TaskService,
) *handler.Server {

	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	resolver := &Resolver{
		userSrv:           userSrv,
		logger:            logger,
		tradingAccountSrv: tradingAccountSrv,
		taskSrv:           taskSrv,
	}

	graphSrv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers:  resolver,
				Directives: generated.DirectiveRoot{},
				Complexity: generated.ComplexityRoot{},
			},
		),
	)
	return graphSrv
}
