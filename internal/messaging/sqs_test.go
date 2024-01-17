package messaging_test

import (
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
	_ "unsafe"
)

//go:linkname publishMessage github.com/ugabiga/falcon/internal/messaging.(*SQSMessageHandler).publishMessage
func publishMessage(handler *messaging.SQSMessageHandler, data interface{}) error

//go:linkname pullToSubscribeMessages github.com/ugabiga/falcon/internal/messaging.(*SQSMessageHandler).pullToSubscribeMessages
func pullToSubscribeMessages(handler *messaging.SQSMessageHandler) error

func TestSQSMessageHandler_PublishAndSubscribeMessage(t *testing.T) {
	tester := app.InitializeTestApplication()
	h := tester.SQSHandler

	//TaskOrderInfoMessage
	msg := messaging.TaskOrderInfoMessage{
		TaskOrderInfo: service.TaskOrderInfo{
			TaskType:         "Grid",
			TaskID:           "task_id",
			TradingAccountID: "trading_account_id",
			UserID:           "user_id",
		},
	}

	if err := publishMessage(h, msg); err != nil {
		t.Fatal(err)
	}

	if err := pullToSubscribeMessages(h); err != nil {
		t.Fatal(err)
	}

}
