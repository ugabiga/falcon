package messaging

import (
	"context"
	"github.com/ugabiga/falcon/internal/messaging/sqs"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
)

type MessageHandler interface {
	Publish() error
	Subscribe() error
}

func NewMessageHandler(
	cfg *config.Config,
	dcaSrv *service.DcaService,
	gridSrv *service.GridService,
	historyService service.TaskHistoryService,
) MessageHandler {

	log.Println("[TEMP] update all task history TTL")
	err := historyService.UpdateAllTaskHistoryTTL(context.Background())
	if err != nil {
		log.Printf("error updating all task history TTL: %v", err)
		return nil
	}

	if cfg.MessagingPlatform != "sqs" {
		panic("invalid messaging platform")
	}

	sqsClient := sqs.NewClient(cfg.SQSQueueURL, cfg.AWSRegion)
	core := sqs.NewMessageCore(
		cfg,
		dcaSrv,
		gridSrv,
		sqsClient,
	)

	switch cfg.SQSSubscriptionType {
	case "local":
		return sqs.NewLocalHandler(
			core,
			sqsClient,
		)
	case "lambda":
		return sqs.NewLambdaHandler(
			core,
		)
	default:
		panic("invalid subscription type")
	}
}
