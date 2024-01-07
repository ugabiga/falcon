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
func (a App) RunPublisher() error {
	return a.messaging.Publish()
}

func (a App) RunSubscriber() error {
	return a.messaging.Subscribe()
}

//func (a App) Worker() error {
//	go a.messaging.Watch()
//	return a.messaging.Listen()
//}
//func (a App) RunLambdaSQS(newMsg messaging.DCAMessage) error {
//	return a.messaging.LambdaSQS(newMsg)
//}
//func (a App) RunLambdaCron() error {
//	return a.messaging.WatchSQS()
//}
