package sqs

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type LambdaHandler struct {
	core *MessageCore
}

func NewLambdaHandler(
	core *MessageCore,
) *LambdaHandler {
	return &LambdaHandler{
		core: core,
	}
}

func (h LambdaHandler) Publish() error {
	lambda.Start(h.handlePublish)
	return nil
}
func (h LambdaHandler) handlePublish() (string, error) {
	if err := h.core.PublishMessages(); err != nil {
		return "FAIL", err
	}

	return "OK", nil
}

func (h LambdaHandler) Subscribe() error {
	lambda.Start(h.handleSubscribe)
	return nil
}
func (h LambdaHandler) handleSubscribe(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	for _, message := range sqsEvent.Records {
		var reqData TaskOrderInfoMessage
		err := json.Unmarshal([]byte(message.Body), &reqData)
		if err != nil {
			log.Println("Error unmarshalling message,", err)
			continue
		}

		if err := h.core.SubscribeMessage(reqData); err != nil {
			log.Println("Error subscribing message,", err)
			continue
		}
	}

	return "OK", nil
}
