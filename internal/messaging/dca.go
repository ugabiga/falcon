package messaging

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

const (
	DCAHandlerName   = "dca_handler"
	DCAIncomingTopic = "dca_incoming_topic"
	DCAOutgoingTopic = "dca_outgoing_topic"
)

type HandlerInfo struct {
	Name          string
	IncomingTopic string
	OutgoingTopic string
	Handler       message.HandlerFunc
}

type DCAMessage struct {
	Msg string
}

type DCAMessageHandler struct {
}

func NewDCAHandler() *DCAMessageHandler {
	return &DCAMessageHandler{}
}

func (h DCAMessageHandler) Handler() HandlerInfo {
	return HandlerInfo{
		Name:          DCAHandlerName,
		IncomingTopic: DCAIncomingTopic,
		OutgoingTopic: DCAOutgoingTopic,
		Handler:       h.Handle,
	}
}

func (h DCAMessageHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	var newMsg DCAMessage
	if err := json.Unmarshal(msg.Payload, &newMsg); err != nil {
		log.Fatalf("Error occurred during unmarshalling. Err: %v", err)
	}

	log.Printf("newMsg: %+v", newMsg)

	return nil, nil
}

func (h DCAMessageHandler) watch() ([]byte, error) {
	data, err := json.Marshal(DCAMessage{Msg: "hello world"})
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (h DCAMessageHandler) Watch(pubSub **gochannel.GoChannel) {
	for {
		log.Printf("Publishing message to topic: %s", DCAIncomingTopic)

		data, err := h.watch()
		if err != nil {
			log.Printf("Error occurred during watching. Err: %v", err)
			continue
		}

		if err := publish(*pubSub, DCAIncomingTopic, data); err != nil {
			log.Printf("Error occurred during publishing. Err: %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}
