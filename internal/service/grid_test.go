package service_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"log"
	"testing"
	"time"
)

type MockTimer struct {
	mock.Mock
}

func NewMockTimer() *MockTimer {
	return &MockTimer{}
}

func (m *MockTimer) NoSeconds() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func TestGridService_GetTarget(t *testing.T) {
	ctx := context.Background()

	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	user := tester.CreateOrGetTestUser(ctx, t)
	tradingAccount := tester.CreateTestTradingAccount(ctx, t, user.ID, model.ExchangeUpbit, "test", "test")

	t.Run("should return tasks", func(t *testing.T) {
		//Create Task
		task, err := tester.TaskSrv.Create(
			ctx,
			user.ID,
			generated.CreateTaskInput{
				TradingAccountID: tradingAccount.ID,
				Currency:         "KRW",
				Size:             0.00001,
				Symbol:           "BTC",
				Hours:            time.Now().Format("18"),
				Days:             "1,2,3,4,5,6,7",
				Type:             model.TaskTypeLongGrid,
				Params:           map[string]interface{}{},
			})
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("NextExecutionTime: %+v", task.NextExecutionTime)

		mt := NewMockTimer()
		mt.On("NoSeconds").Return(task.NextExecutionTime)

		target, err := tester.GridSrv.GetTarget(mt)
		if err != nil {
			t.Error(err)
		}
		log.Printf("target: %+v", target)

		if len(target) < 1 {
			t.Errorf("expected len(target) > 0, got %d", len(target))
		}
	})
}

func TestGridService_Order(t *testing.T) {
	ctx := context.Background()

	tester := app.InitializeTestApplication()
	cfg := tester.Cfg
	tester.ResetTables(t)
	user := tester.CreateOrGetTestUser(ctx, t)

	t.Run("should make upbit order", func(t *testing.T) {
		tradingAccount := tester.CreateTestTradingAccount(
			ctx,
			t,
			user.ID,
			model.ExchangeUpbit,
			cfg.TestUpbitKey,
			cfg.TestUpbitSecret,
		)

		gridTaskParams := model.TaskGridParams{
			GapPercent: 3,
			Quantity:   2,
		}

		task, err := tester.TaskSrv.Create(
			ctx,
			user.ID,
			generated.CreateTaskInput{
				TradingAccountID: tradingAccount.ID,
				Currency:         "KRW",
				Size:             0.0001,
				Symbol:           "BTC",
				Hours:            time.Now().Format("18"),
				Days:             "1,2,3,4,5,6,7",
				Type:             model.TaskTypeLongGrid,
				Params:           gridTaskParams.ToParams(),
			})
		if err != nil {
			t.Fatal(err)
		}

		mt := NewMockTimer()
		mt.On("NoSeconds").Return(task.NextExecutionTime)

		target, err := tester.GridSrv.GetTarget(mt)
		if err != nil {
			t.Error(err)
		}

		if len(target) < 1 {
			t.Errorf("expected len(target) > 0, got %d", len(target))
		}

		err = tester.GridSrv.Order(target[0])
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should make binance order", func(t *testing.T) {
		tradingAccount := tester.CreateTestTradingAccount(
			ctx,
			t,
			user.ID,
			model.ExchangeBinanceFutures,
			cfg.TestBinanceKey,
			cfg.TestBinanceSecret,
		)

		gridTaskParams := model.TaskGridParams{
			GapPercent: 5,
			Quantity:   1,
		}

		task, err := tester.TaskSrv.Create(
			ctx,
			user.ID,
			generated.CreateTaskInput{
				TradingAccountID: tradingAccount.ID,
				Currency:         "USDT",
				Size:             0.003,
				Symbol:           "BTC",
				Hours:            time.Now().Format("18"),
				Days:             "1,2,3,4,5,6,7",
				Type:             model.TaskTypeLongGrid,
				Params:           gridTaskParams.ToParams(),
			})
		if err != nil {
			t.Fatal(err)
		}

		mt := NewMockTimer()
		mt.On("NoSeconds").Return(task.NextExecutionTime)

		target, err := tester.GridSrv.GetTarget(mt)
		if err != nil {
			t.Error(err)
		}

		if len(target) < 1 {
			t.Errorf("expected len(target) > 0, got %d", len(target))
		}

		err = tester.GridSrv.Order(target[0])
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("should make binance spot order", func(t *testing.T) {
		tradingAccount := tester.CreateTestTradingAccount(
			ctx,
			t,
			user.ID,
			model.ExchangeBinanceSpot,
			cfg.TestBinanceKey,
			cfg.TestBinanceSecret,
		)

		gridTaskParams := model.TaskGridParams{
			GapPercent: 5,
			Quantity:   2,
		}

		task, err := tester.TaskSrv.Create(
			ctx,
			user.ID,
			generated.CreateTaskInput{
				TradingAccountID: tradingAccount.ID,
				Currency:         "USDT",
				Size:             0.0002,
				Symbol:           "BTC",
				Hours:            time.Now().Format("15"),
				Days:             "1,2,3,4,5,6,7",
				Type:             model.TaskTypeLongGrid,
				Params:           gridTaskParams.ToParams(),
			})
		if err != nil {
			t.Fatal(err)
		}

		mt := NewMockTimer()
		mt.On("NoSeconds").Return(task.NextExecutionTime)

		target, err := tester.GridSrv.GetTarget(mt)
		if err != nil {
			t.Error(err)
		}

		if len(target) < 1 {
			t.Errorf("expected len(target) > 0, got %d", len(target))
		}

		err = tester.GridSrv.Order(target[0])
		if err != nil {
			t.Error(err)
		}

	})

}
