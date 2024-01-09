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

func TestRepository_Migration(t *testing.T) {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	repo := tester.Repository

	ipAddress := "3.39.156.133"
	ctx := context.Background()

	//Create static ip
	_, err := repo.CreateStaticIP(ctx, model.StaticIP{
		IPAddress:      ipAddress,
		IPAvailability: true,
		IPUsageCount:   0,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepository(t *testing.T) {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	repo := tester.Repository

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

	targetTask.IsActive = false
	_, err = repo.UpdateTask(ctx, *targetTask)
	if err != nil {
		t.Fatal(err)
	}

	tasksByNextExecutionTime, err := repo.GetTasksByActiveNextExecutionTime(ctx, targetTask.NextExecutionTime)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("tasksByNextExecutionTime: %+v", debug.ToJSONStr(tasksByNextExecutionTime))

	if len(tasksByNextExecutionTime) != 10-1 {
		t.Fatal("tasks count is not equal")
	}
}

func TestDynamoRepository_StaticIP(t *testing.T) {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	repo := tester.Repository
	ipAddress := "192.168.0.1"
	ctx := context.Background()

	//Create static ip
	staticIP, err := repo.CreateStaticIP(ctx, model.StaticIP{
		IPAddress:      ipAddress,
		IPAvailability: true,
		IPUsageCount:   0,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should count up usage", func(t *testing.T) {
		//Count up usage
		err = repo.CountUpStaticIPUsage(ctx, staticIP.ID)
		if err != nil {
			t.Fatal(err)
		}

		//Retrieve static ip
		retrieveStaticIP, err := repo.GetStaticIP(ctx, staticIP.ID)
		if err != nil {
			t.Fatal(err)
		}

		log.Printf("retrieveStaticIP: %+v", debug.ToJSONStr(retrieveStaticIP))

		if retrieveStaticIP.IPUsageCount != 1 {
			t.Fatal("usage count is not equal")
		}
	})

	t.Run("should count down usage", func(t *testing.T) {
		//Count down usage
		err = repo.CountDownStaticIPUsage(ctx, staticIP.ID)
		if err != nil {
			t.Fatal(err)
		}

		//Retrieve static ip
		retrieveStaticIP, err := repo.GetStaticIP(ctx, staticIP.ID)
		if err != nil {
			t.Fatal(err)
		}

		log.Printf("retrieveStaticIP: %+v", debug.ToJSONStr(retrieveStaticIP))

		if retrieveStaticIP.IPUsageCount != 0 {
			t.Fatal("usage count is not equal")
		}
	})

	t.Run("should find available ip", func(t *testing.T) {
		staticIPByAvailability, err := repo.GetStaticIPByAvailability(ctx)
		if err != nil {
			t.Fatal(err)
		}

		log.Printf("staticIPByAvailability: %+v", debug.ToJSONStr(staticIPByAvailability))

		if staticIPByAvailability.IPAddress != ipAddress {
			t.Fatal("ip address is not equal")
		}

	})
}
