package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"log"
	"testing"
	"time"
)

func TestDcaService_GetTarget(t *testing.T) {
	tester := Initialize(t)
	srv := tester.DcaSrv

	t.Run("should return tasks", func(t *testing.T) {
		//Create User
		ctx := context.Background()

		authentication, user, err := tester.AuthenticationSrv.SignUp(
			ctx,
			"google",
			uuid.New().String(),
			uuid.New().String(),
			"new_user",
		)
		if err != nil {
			t.Fatal(err)
		}

		if authentication == nil {
			t.Fatal("authentication is nil")
		}

		if user == nil {
			t.Fatal("user is nil")
		}

		//Create TradingAccount
		tradingAccount, err := tester.TradingAccountSrv.Create(
			ctx,
			user.ID,
			"test",
			"binance",
			"test",
			"test",
			"test",
		)
		if err != nil {
			t.Fatal(err)
		}

		if tradingAccount == nil {
			t.Fatal("tradingAccount is nil")
		}

		//Create Task
		task, err := tester.TaskSrv.Create(
			ctx,
			user.ID,
			generated.CreateTaskInput{
				TradingAccountID: tradingAccount.ID,
				Currency:         "KRW",
				Size:             1,
				Symbol:           "BTC",
				Hours:            time.Now().Format("18"),
				Days:             "1,2,3,4,5,6,7",
				Type:             "test",
				Params:           map[string]interface{}{},
			})
		if err != nil {
			t.Fatal(err)
		}

		if task == nil {
			t.Fatal("task is nil")
		}
		log.Printf("NextExecutionTime: %+v", task.NextExecutionTime)

		//Get Target
		target, err := srv.GetTarget()
		if err != nil {
			t.Fatal(err)
		}

		if target == nil {
			t.Fatal("target is nil")
		}

		log.Printf("Target: %+v", target)

	})
}
