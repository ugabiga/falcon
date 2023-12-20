package app

import (
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/handler"
	"github.com/ugabiga/falcon/internal/server"
	"github.com/ugabiga/falcon/internal/service"
	"go.uber.org/fx"
)

func provider() fx.Option {
	return fx.Provide(
		config.NewConfig,

		// Client
		client.NewEntClient,

		// Service
		service.NewJWTService,
		service.NewUserService,
		service.NewAuthenticationService,
		service.NewTradingAccountService,

		// Handler
		handler.NewAuthenticationHandler,
		handler.NewUserHandler,

		// Server
		server.NewServer,

		// App
		NewApp,
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
