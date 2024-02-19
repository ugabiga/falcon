package app

import (
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/server"
	"github.com/ugabiga/falcon/pkg/config"
)

type App struct {
	cfg       *config.Config
	server    *server.Server
	messaging messaging.MessageHandler
}

func NewApp(
	cfg *config.Config,
	server *server.Server,
	messaging messaging.MessageHandler,
) App {
	return App{
		cfg:       cfg,
		server:    server,
		messaging: messaging,
	}
}

func (a App) RunLambdaServer() error {
	return a.server.RunLambda()
}
func (a App) RunServer() error {
	return a.server.Run()
}
func (a App) RunCron() error {
	return a.messaging.Publish()
}
func (a App) RunWorker() error {
	return a.messaging.Subscribe()
}
