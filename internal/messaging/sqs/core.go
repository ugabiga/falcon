package sqs

import (
	"context"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
)

type MessageCore struct {
	sqsClient      *Client
	dcaSrv         *service.DcaService
	gridSrv        *service.GridService
	taskHistorySrv *service.TaskHistoryService
}

func NewMessageCore(
	cfg *config.Config,
	dcaSrv *service.DcaService,
	gridSrv *service.GridService,
	sqsClient *Client,
	taskHistorySrv *service.TaskHistoryService,
) *MessageCore {
	return &MessageCore{
		dcaSrv:         dcaSrv,
		gridSrv:        gridSrv,
		sqsClient:      sqsClient,
		taskHistorySrv: taskHistorySrv,
	}
}

func (c MessageCore) PublishMessages() error {
	if err := c.publishDCAMessages(); err != nil {
		log.Printf("Error occurred during publishing DCA messages. Err: %v", err)
	}

	if err := c.publishLongGridMessages(); err != nil {
		log.Printf("Error occurred during publishing LongGrid messages. Err: %v", err)
	}

	msg := TaskOrderInfoMessage{
		TaskOrderInfo: service.TaskOrderInfo{
			TaskType:         "migration",
			TaskID:           "migration",
			TradingAccountID: "migration",
			UserID:           "migration",
		},
	}
	if _, err := c.sqsClient.Publish(msg); err != nil {
		log.Printf("Error occurred during publishing. Err: %v", err)
	}

	return nil
}
func (c MessageCore) publishLongGridMessages() error {
	messages, err := c.gridSrv.GetTarget(nil)
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		log.Printf("No messages to publish")
		return nil
	}

	var gridMessages []TaskOrderInfoMessage
	for _, msg := range messages {
		gridMessages = append(gridMessages, TaskOrderInfoMessage{
			TaskOrderInfo: msg,
		})
	}

	for _, msg := range gridMessages {
		if _, err := c.sqsClient.Publish(msg); err != nil {
			log.Printf("Error occurred during publishing. Err: %v", err)
			continue
		}
	}

	return nil
}

func (c MessageCore) publishDCAMessages() error {
	messages, err := c.dcaSrv.GetTarget()
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		log.Printf("No messages to publish")
		return nil
	}

	var dcaMessages []TaskOrderInfoMessage
	for _, msg := range messages {
		dcaMessages = append(dcaMessages, TaskOrderInfoMessage{
			TaskOrderInfo: msg,
		})
	}

	for _, msg := range dcaMessages {
		if _, err := c.sqsClient.Publish(msg); err != nil {
			log.Printf("Error occurred during publishing. Err: %v", err)
			continue
		}
	}

	return nil
}

func (c MessageCore) SubscribeMessage(reqMsg TaskOrderInfoMessage) error {
	switch reqMsg.TaskOrderInfo.TaskType {
	case model.TaskTypeDCA:
		if err := c.dcaSrv.Order(reqMsg.TaskOrderInfo); err != nil {
			log.Printf("Error occurred during order. Err: %v", err)
			return err
		}
	case model.TaskTypeLongGrid:
		if err := c.gridSrv.Order(reqMsg.TaskOrderInfo); err != nil {
			log.Printf("Error occurred during order. Err: %v", err)
			return err
		}
		return nil
	case "migration":
		log.Println("[TEMP] update all task history TTL")
		err := c.taskHistorySrv.UpdateAllTaskHistoryTTL(context.Background())
		if err != nil {
			log.Printf("error updating all task history TTL: %v", err)
			return nil
		}
		log.Printf("Migration task received")
		return nil
	default:
		log.Printf("Unknown task type: %s", reqMsg.TaskOrderInfo.TaskType)
		return nil
	}

	return nil
}
