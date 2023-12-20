package service_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
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
			uuid.New().String(),
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
			uuid.New().String(),
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
			uuid.New().String(),
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

func TestTradingAccountService_GetByID(t *testing.T) {
	ctx := context.Background()
	user := prepareUser(t)
	srv := initTradingAccountService()

	t.Run("should return a trading account", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"USD",
			uuid.New().String(),
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

		a, err = srv.GetByID(ctx, user.ID, a.ID)
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
}

func TestTradingAccountService_Get(t *testing.T) {
	ctx := context.Background()
	user := prepareUser(t)
	srv := initTradingAccountService()

	t.Run("should return trading accounts", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"USD",
			uuid.New().String(),
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

		accounts, err := srv.Get(ctx, user.ID)
		if err != nil {
			t.Fatal(err)
		}

		if len(accounts) == 0 {
			t.Fatal("accounts is empty")
		}
	})
}

func TestTradingAccountService_Update(t *testing.T) {
	ctx := context.Background()
	user := prepareUser(t)
	srv := initTradingAccountService()

	t.Run("should update a trading account", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"USD",
			uuid.New().String(),
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

		err = srv.Update(
			ctx,
			a.ID,
			user.ID,
			"binance",
			"USD",
			uuid.New().String(),
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
}

func TestTradingAccountService_Delete(t *testing.T) {
	ctx := context.Background()
	user := prepareUser(t)
	srv := initTradingAccountService()

	t.Run("should delete a trading account", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			user.ID,
			"binance",
			"USD",
			uuid.New().String(),
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

		err = srv.Delete(ctx, user.ID, a.ID)
		if err != nil {
			t.Fatal(err)
		}
	})
}
