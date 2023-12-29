package messaging

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"time"
)

type Messaging struct {
	pubSub *gochannel.GoChannel
	logger watermill.LoggerAdapter
	router *message.Router

	dcaMessageHandler *DCAMessageHandler
}

func NewMessaging(
	dcaMessageHandler *DCAMessageHandler,
) (*Messaging, error) {
	return &Messaging{
		dcaMessageHandler: dcaMessageHandler,
	}, nil
}

func (m *Messaging) route() error {
	router, err := message.NewRouter(message.RouterConfig{}, m.logger)
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
			Logger:          m.logger,
		}.Middleware,
		middleware.Recoverer,
	)

	// Handler
	m.addHandler(router, m.dcaMessageHandler.Handler())
	m.router = router
	return nil
}

func (m *Messaging) Watch() {
	go m.dcaMessageHandler.Watch(&m.pubSub)
}

func (m *Messaging) Listen() error {
	ctx := context.Background()

	m.logger = watermill.NewStdLogger(false, false)
	m.pubSub = gochannel.NewGoChannel(gochannel.Config{}, m.logger)

	if err := m.route(); err != nil {
		return err
	}

	if err := m.router.Run(ctx); err != nil {
		return err
	}
	return nil
}

func (m *Messaging) addHandler(router *message.Router, handlerInfo HandlerInfo) {
	router.AddHandler(
		handlerInfo.Name,
		handlerInfo.IncomingTopic,
		m.pubSub,
		handlerInfo.OutgoingTopic,
		m.pubSub,
		handlerInfo.Handler,
	)
}

func publish(pubSub *gochannel.GoChannel, topic string, data []byte) error {
	newMsg := message.NewMessage(watermill.NewUUID(), data)

	if err := pubSub.Publish(topic, newMsg); err != nil {
		return err
	}

	return nil
}
