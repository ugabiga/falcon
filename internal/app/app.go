package app

import (
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/server"
	"github.com/ugabiga/falcon/pkg/config"
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
func (a App) Worker() error {
	go a.messaging.Watch()
	return a.messaging.Listen()
	//go func() {
	//	err := a.messaging.Listen()
	//	if err != nil {
	//		log.Fatalf("Error occurred during listening. Err: %v", err)
	//	}
	//}()
	//
	//a.messaging.Watch()
	//
	//time.Sleep(60 * time.Second)
	//
	//return nil
}
func (a App) RunLambdaServer() error {
	return a.server.RunLambda()
}
func (a App) RunLambdaSQS(newMsg messaging.DCAMessage) error {
	return a.messaging.LambdaSQS(newMsg)
}
func (a App) RunLambdaCron() error {
	return a.messaging.WatchSQS()
}
