package messaging

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/service"
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
	TaskOrderInfo service.TaskOrderInfo
}

type DCAMessageHandler struct {
	dcaSrv *service.DcaService
}

func NewDCAHandler(
	dcaSrv *service.DcaService,
) *DCAMessageHandler {
	return &DCAMessageHandler{
		dcaSrv: dcaSrv,
	}
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

	log.Printf("newMsg: %+v", debug.ToJSONStr(newMsg))

	if err := h.dcaSrv.Order(newMsg.TaskOrderInfo); err != nil {
		log.Printf("Error occurred during order. Err: %v", err)
		return nil, err
	}

	return nil, nil
}

func (h DCAMessageHandler) dcaMessages() ([]DCAMessage, error) {
	orders, err := h.dcaSrv.GetTarget()
	if err != nil {
		return nil, err
	}

	var dcaMessages []DCAMessage
	for _, order := range orders {
		dcaMessages = append(dcaMessages, DCAMessage{
			TaskOrderInfo: order,
		})
	}

	return dcaMessages, nil
}
func (h DCAMessageHandler) Watch(pubSub **gochannel.GoChannel) {
	for {
		log.Printf("Watching for DCA messages...")
		messages, err := h.dcaMessages()
		if err != nil {
			log.Printf("Error occurred during watching. Err: %v", err)
			continue
		}

		for _, msg := range messages {
			data, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error occurred during marshalling. Err: %v", err)
				continue
			}

			if err := publish(*pubSub, DCAIncomingTopic, data); err != nil {
				log.Printf("Error occurred during publishing. Err: %v", err)
				continue
			}
			if err != nil {
				log.Printf("Error occurred during publishing. Err: %v", err)
			}
		}

		time.Sleep(time.Minute)
	}
}
