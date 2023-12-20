package service_test

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
)

func initTradingAccountService() *service.TradingAccountService {
	cfg := &config.Config{
		DBDriverName: "sqlite3",
		DBSource:     "file:ent?mode=memory&cache=shared&_fk=1",
	}
	entClient := client.NewEntClient(cfg)

	return service.NewTradingAccountService(entClient)
}

func TestTradingAccountService_Create(t *testing.T) {
	ctx := context.Background()
	user := prepareUser(t)
	srv := initTradingAccountService()

	t.Run("should create a trading account", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"USD",
			"identifier",
			"credential",
			"",
		)
		if err != nil {
			t.Fatal(err)
		}

		if a == nil {
			t.Fatal("a is nil")
		}

		if a.ID == 0 {
			t.Fatal("a.ID is 0")
		}
	})

	t.Run("should return an error if exchange is wrong", func(t *testing.T) {
		t.Parallel()

		_, err := srv.Create(
			ctx,
			user.ID,
			"wrong_exchange",
			"USD",
			"identifier",
			"credential",
			"",
		)
		if err == nil {
			t.Fatal("err is nil")
		}

		if !errors.Is(err, service.ErrWrongExchange) {
			t.Fatal("err is not ErrWrongExchange")
		}
	})

	t.Run("should return an error if currency is wrong", func(t *testing.T) {
		t.Parallel()

		_, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"wrong_currency",
			"identifier",
			"credential",
			"",
		)
		if err == nil {
			t.Fatal("err is nil")
		}

		if !errors.Is(err, service.ErrWrongCurrency) {
			t.Fatal("err is not ErrWrongCurrency")
		}
	})
}
