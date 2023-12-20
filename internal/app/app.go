package app

import (
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/server"
)

type App struct {
	cfg    *config.Config
	server *server.Server
}

func NewApp(
	cfg *config.Config,
	server *server.Server,
) App {
	return App{
		cfg:    cfg,
		server: server,
	}
}

func (a App) RunServer() error {
	return a.server.Run()
}
