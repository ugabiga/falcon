package app

import (
	"github.com/ugabiga/falcon/internal/config"
	"go.uber.org/fx"
)

func provider() fx.Option {
	return fx.Provide(
		config.NewConfig,
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
