package app

import (
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/graph/helper"
	"github.com/ugabiga/falcon/internal/graph/resolvers"
	"github.com/ugabiga/falcon/internal/handler"
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/server"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"go.uber.org/fx"
)

func provider() fx.Option {
	return fx.Provide(
		config.NewConfig,

		// Client
		client.NewEntClient,

		// Service
		service.NewUserService,
		service.NewAuthenticationService,
		service.NewTradingAccountService,
		service.NewTaskService,
		service.NewTaskHistoryService,
		service.NewDcaService,

		// Handler
		handler.NewAuthenticationHandler,
		handler.NewUserHandler,
		handler.NewHomeHandler,
		handler.NewErrorHandler,
		handler.NewTradingAccountHandler,
		handler.NewTaskHandler,

		// GraphQL
		helper.NewNodeResolver,
		resolvers.NewResolver,

		// Server
		server.NewServer,

		//Messaging Handler
		messaging.NewDCAHandler,

		//Messaging
		messaging.NewMessaging,

		//Others
		encryption.NewEncryption,

		// App
		NewApp,
	)
}

func InitializeApplication() App {
	var newApp App

	fx.New(
		//fx.NopLogger,
		provider(),
		fx.Invoke(func(lifecycle fx.Lifecycle, app App) {
			newApp = app
		}),
	)

	return newApp
}
