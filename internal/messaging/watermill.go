package messaging

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/service"
	"log"
	"time"
)

type WatermillMessageHandler struct {
	dcaSrv *service.DcaService
	logger watermill.LoggerAdapter
	pubSub *gochannel.GoChannel
	router *message.Router
}

func NewWatermillMessageHandler(
	dcaSrv *service.DcaService,
) *WatermillMessageHandler {
	logger := watermill.NewStdLogger(false, false)

	return &WatermillMessageHandler{
		dcaSrv: dcaSrv,
		logger: logger,
		pubSub: gochannel.NewGoChannel(gochannel.Config{}, logger),
	}
}

func (h *WatermillMessageHandler) Publish() error {
	for {
		log.Printf("Watching for DCA messages...")
		messages, err := h.dcaMessages()
		if err != nil {
			log.Printf("Error occurred during watching. Err: %v", err)
			return err
		}

		for _, msg := range messages {
			data, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error occurred during marshalling. Err: %v", err)
				continue
			}

			if err := publish(h.pubSub, DCAIncomingTopic, data); err != nil {
				log.Printf("Error occurred during publishing. Err: %v", err)
				continue
			}
			if err != nil {
				log.Printf("Error occurred during publishing. Err: %v", err)
			}
		}

		log.Printf("Watching for DCA messages... Done")
		time.Sleep(30 * time.Second)
	}
}

func (h *WatermillMessageHandler) Subscribe() error {
	ctx := context.Background()

	if err := h.route(); err != nil {
		return err
	}

	if err := h.router.Run(ctx); err != nil {
		return err
	}
	return nil
}

func (h *WatermillMessageHandler) route() error {
	router, err := message.NewRouter(message.RouterConfig{}, h.logger)
	if err != nil {
		return err
	}

	// Middleware
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      5,
			InitialInterval: time.Millisecond * 10,
			Logger:          h.logger,
		}.Middleware,
		middleware.Recoverer,
	)

	// Handler
	h.addHandler(router, HandlerInfo{
		Name:          DCAHandlerName,
		IncomingTopic: DCAIncomingTopic,
		OutgoingTopic: DCAOutgoingTopic,
		Handler:       h.subscribe,
	})
	h.router = router
	return nil
}

func (h *WatermillMessageHandler) addHandler(router *message.Router, handlerInfo HandlerInfo) {
	router.AddHandler(
		handlerInfo.Name,
		handlerInfo.IncomingTopic,
		h.pubSub,
		handlerInfo.OutgoingTopic,
		h.pubSub,
		handlerInfo.Handler,
	)
}

func (h *WatermillMessageHandler) subscribe(msg *message.Message) ([]*message.Message, error) {
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

func (h *WatermillMessageHandler) dcaMessages() ([]DCAMessage, error) {
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

func publish(pubSub *gochannel.GoChannel, topic string, data []byte) error {
	newMsg := message.NewMessage(watermill.NewUUID(), data)

	if err := pubSub.Publish(topic, newMsg); err != nil {
		return err
	}

	return nil
}
