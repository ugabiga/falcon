package messaging

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ugabiga/falcon/internal/service"
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
