package sqs

import (
	"encoding/json"
	"github.com/AlekSi/pointer"
	"log"
)

type LocalHandler struct {
	core      *MessageCore
	sqsClient *Client
}

func NewLocalHandler(
	core *MessageCore,
	sqsClient *Client,
) *LocalHandler {
	return &LocalHandler{
		core:      core,
		sqsClient: sqsClient,
	}
}

func (h LocalHandler) Publish() error {
	return h.core.PublishMessages()
}

func (h LocalHandler) Subscribe() error {
	messages, err := h.sqsClient.Pull()
	if err != nil {
		return err
	}

	if len(messages.Messages) == 0 {
		return nil
	}
	log.Printf("Received %d messages", len(messages.Messages))

	for _, msg := range messages.Messages {
		if msg.Body == nil {
			log.Println("Error getting message body, body is nil")
			continue
		}

		var reqMsg TaskOrderInfoMessage
		if err := json.Unmarshal([]byte(pointer.GetString(msg.Body)), &reqMsg); err != nil {
			log.Println("Error unmarshalling message,", err)
			continue
		}

		if err := h.core.SubscribeMessage(reqMsg); err != nil {
			log.Println("Error subscribing message,", err)
		}

		if err := h.sqsClient.Delete(msg); err != nil {
			log.Println("Error deleting message,", err)
		}
	}

	return nil
}
