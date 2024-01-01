package app

import (
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/server"
)

type App struct {
	cfg       *config.Config
	server    *server.Server
	messaging *messaging.Messaging
}

func NewApp(
	cfg *config.Config,
	server *server.Server,
	messaging *messaging.Messaging,
) App {
	return App{
		cfg:       cfg,
		server:    server,
		messaging: messaging,
	}
}

func (a App) RunServer() error {
	return a.server.Run()
}
func (a App) WatchMessageAndListen() error {
	go a.messaging.Watch()
	return a.messaging.Listen()
}
func (a App) RunLambdaServer() error {
	return a.server.RunLambda()
}
