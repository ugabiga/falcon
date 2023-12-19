package app

import (
	"context"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/service"
	"log"
)

type App struct {
	cfg     *config.Config
	userSrv *service.UserService
}

func NewApp(
	cfg *config.Config,
	userSrv *service.UserService,
) App {
	return App{
		cfg:     cfg,
		userSrv: userSrv,
	}
}

func (a App) RunServer() error {
	log.Printf("Run server")

	ctx := context.Background()

	if err := a.userSrv.GetUser(ctx); err != nil {
		return err
	}

	return nil
}
