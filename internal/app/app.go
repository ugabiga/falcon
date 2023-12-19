package app

import (
	"github.com/ugabiga/falcon/internal/config"
	"log"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) App {
	return App{
		cfg: cfg,
	}
}

func (a App) RunServer() error {
	log.Printf("Run server")
	log.Printf(a.cfg.DBDriverName)
	return nil
}
