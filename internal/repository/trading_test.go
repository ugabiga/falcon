package repository_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/model"
	"log"
	"testing"
	"time"
)

func TestTradingRepository_Migration(t *testing.T) {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
}

func TestTradingRepository(t *testing.T) {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	repo := tester.TradingRepository

	ctx := context.Background()

	userID := uuid.New().String()

	tradingAccount, err := repo.CreateTradingAccount(
		ctx,
		model.TradingAccount{
			UserID:   userID,
			Name:     "Default",
			Exchange: "binance",
			IP:       "192.168.0.1",
			Key:      "key",
			Secret:   "secret",
			Phrase:   "phrase",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if tradingAccount.ID == "" {
		t.Fatal("ID is empty")
	}

	tradingAccountID := tradingAccount.ID

	retrieveTradingAccount, err := repo.GetTradingAccount(ctx, userID, tradingAccountID)
	if err != nil {
		t.Fatal(err)
	}

	if retrieveTradingAccount.ID != tradingAccountID {
		t.Fatal("ID is not equal")
	}

	log.Printf("tradingAccount: %+v", debug.ToJSONStr(tradingAccount))

	var targetTask *model.Task
	for i := 0; i < 10; i++ {
		createdTask, err := repo.CreateTask(
			ctx,
			model.Task{
				TradingAccountID:  tradingAccountID,
				UserID:            tradingAccount.UserID,
				Currency:          "USD",
				Symbol:            "BTC",
				Size:              0.02,
				Cron:              "0 0 0 1 1 *",
				NextExecutionTime: time.Now(),
				IsActive:          true,
				Type:              "dca",
			})
		if err != nil {
			t.Fatal(err)
		}

		if createdTask.ID == "" {
			t.Fatal("ID is empty")
		}

		targetTask = createdTask
	}

	log.Printf("task: %+v", debug.ToJSONStr(targetTask))

	retrieveTask, err := repo.GetTask(ctx, targetTask.TradingAccountID, targetTask.ID)
	if err != nil {
		t.Fatal(err)
	}

	if retrieveTask.ID != targetTask.ID {
		t.Fatal("ID is not equal")
	}

	log.Printf("retrieveTask: %+v", debug.ToJSONStr(retrieveTask))

	tasks, err := repo.GetTasksByTradingAccountID(ctx, tradingAccountID)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 10 {
		t.Fatal("tasks count is not equal")
	}
	log.Printf("tasks: %+v", debug.ToJSONStr(tasks))

	taskCount, err := repo.CountTasksByTradingID(ctx, tradingAccountID)
	if err != nil {
		t.Fatal(err)
	}

	if taskCount != 10 {
		t.Fatal("task count is not equal")
	}

	tasksByNextExecutionTime, err := repo.GetTasksByNextExecutionTime(ctx, targetTask.NextExecutionTime)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("tasksByNextExecutionTime: %+v", debug.ToJSONStr(tasksByNextExecutionTime))

	if len(tasksByNextExecutionTime) != 10 {
		t.Fatal("tasks count is not equal")
	}
}
