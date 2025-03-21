package app

import (
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/graph/resolvers"
	"github.com/ugabiga/falcon/internal/handler"
	v1 "github.com/ugabiga/falcon/internal/handler/v1"
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/migration"
	"github.com/ugabiga/falcon/internal/repository"
	"github.com/ugabiga/falcon/internal/server"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"go.uber.org/fx"
)

func provider() fx.Option {
	return fx.Provide(
		config.NewConfig,

		// Client
		client.NewDynamoClient,

		// Migration
		migration.NewMigration,

		// Repository
		repository.NewDynamoRepository,

		// Service
		service.NewUserService,
		service.NewAuthenticationService,
		service.NewTradingAccountService,
		service.NewTaskService,
		service.NewTaskHistoryService,
		service.NewDcaService,
		service.NewGridService,
		service.NewMigrationService,

		// Handler
		handler.NewHomeHandler,
		handler.NewAuthenticationHandler,
		handler.NewErrorHandler,

		// Handler V1
		v1.NewUserHandler,
		v1.NewTaskHandler,

		// GraphQL
		resolvers.NewResolver,

		// Server
		server.NewServer,

		//MessagingPlatform Handler
		messaging.NewMessageHandler,

		//Others
		encryption.NewEncryption,

		// App
		NewApp,

		//Tester
		NewTester,
	)
}

func InitializeApplication() App {
	var newApp App

	fx.New(
		fx.NopLogger,
		provider(),
		fx.Invoke(func(lifecycle fx.Lifecycle, app App) {
			newApp = app
		}),
	)

	return newApp
}

func InitializeTestApplication() Tester {
	var newTester Tester

	fx.New(
		fx.NopLogger,
		provider(),
		fx.Invoke(func(lifecycle fx.Lifecycle, tester Tester) {
			newTester = tester
		}),
	)

	return newTester
}
