package client

import (
	"context"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
	"os"
)

func NewEntClient(cfg *config.Config) *ent.Client {
	client, err := ent.Open(cfg.DBDriverName, cfg.DBSource)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func CloseEntClient(entClient *ent.Client) error {
	if err := entClient.Close(); err != nil {
		return err
	}

	return nil
}

func recodeSchemaChanges(client *ent.Client, ctx context.Context) {
	// Create a file for saving schema changes.
	f, err := os.Create("migrate.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close migrate file: %v", err)
		}
	}(f)

	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
