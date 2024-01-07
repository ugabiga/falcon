package service_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
)

func initTradingAccountService() *service.TradingAccountService {
	//cfg := &config.Config{
	//	DBDriverName: "sqlite3",
	//	DBSource:     "file:ent?mode=memory&cache=shared&_fk=1",
	//}
	//entClient := client.NewEntClient(cfg)
	//encryp := encryption.NewEncryption(cfg)
	//
	//return service.NewTradingAccountService(entClient, encryp)
	return nil
}

func TestTradingAccountService_Create(t *testing.T) {
	tester := Initialize(t)
	srv := tester.TradingAccountSrv
	targetUser := prepareUser(t, tester.AuthenticationSrv)
	ctx := context.Background()

	t.Run("should create a trading account", func(t *testing.T) {
		t.Parallel()

		a, err := srv.Create(
			ctx,
			targetUser.ID,
			"binance",
			"upbit",
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

		if a.ID == "" {
			t.Fatal("a.ID is empty")
		}
	})

	t.Run("should return an error if exchange is wrong", func(t *testing.T) {
		t.Parallel()

		_, err := srv.Create(
			ctx,
			targetUser.ID,
			"wrong_exchange",
			"_binance_",
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
}
