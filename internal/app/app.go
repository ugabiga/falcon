package app

import (
	"context"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/ent"
	"log"
)

type App struct {
	cfg       *config.Config
	entClient *ent.Client
}

func NewApp(
	cfg *config.Config,
	entClient *ent.Client,
) App {
	return App{
		cfg:       cfg,
		entClient: entClient,
	}
}

func (a App) RunServer() error {
	log.Printf("Run server")
	ctx := context.Background()
	u, err := a.entClient.User.Query().First(ctx)
	if err != nil {
		return err
	}

	log.Printf("User: %v", u.Name)

	return nil
}
