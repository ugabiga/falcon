package service_test

import (
	"context"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/model"
	"testing"
)

func TestTaskService_MigrateGridParams(t *testing.T) {
	ctx := context.Background()
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)

	user := tester.CreateOrGetTestUser(ctx, t)
	tradingAccount := tester.CreateTestTradingAccount(ctx, t, user.ID, model.ExchangeUpbit, "test", "test")
	_ = tester.CreateTestTasks(ctx, t, tradingAccount, user.ID)

	//log.Printf("Previous tasks: %+v", debug.ToJSONStr(tasks))

	t.Run("should migrate grid params", func(t *testing.T) {
		err := tester.TaskSrv.MigrateGridParams(ctx)
		if err != nil {
			t.Fatal(err)
		}

		migratedTasks, err := tester.TaskSrv.GetByTradingAccount(ctx, tradingAccount.ID)
		if err != nil {
			t.Fatal(err)
		}

		//log.Printf("Migrated tasks: %+v", debug.ToJSONStr(migratedTasks))

		for _, task := range migratedTasks {
			if task.Type == model.TaskTypeLongGrid {
				if task.Params == nil {
					t.Errorf("Task params should not be nil")
				}

				if task.Params["delete_previous_orders"] == nil {
					t.Errorf("Task params should have delete_previous_orders")
				}

				paramsV2, err := task.GridParamsV2()
				if err != nil {
					t.Errorf("Error while parse grid params v2")
				}

				if paramsV2.DeletePreviousOrders != true {
					t.Errorf("Task params should set should delete previous orders to true")
				}
			}
		}
	})

}
