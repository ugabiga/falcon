package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
	"modernc.org/sqlite"

	"os"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		if err := conn.Close(); err != nil {
			return nil, err
		}
		return nil, errors.New("failed to enable foreign keys")
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

func NewEntClient(cfg *config.Config) *ent.Client {
	var client *ent.Client
	var err error

	if cfg.DBDriverName == "sqlite3" {
		db, err := sql.Open(cfg.DBDriverName, cfg.DBSource)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}

		drv := entsql.OpenDB(cfg.DBDriverName, db)
		client = ent.NewClient(ent.Driver(drv))

	} else {
		client, err = ent.Open(cfg.DBDriverName, cfg.DBSource)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
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
