package messaging

import (
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
)

type MessageHandler interface {
	Publish() error
	Subscribe() error
}

func NewMessageHandler(
	cfg *config.Config,
	dcaSrv *service.DcaService,
) MessageHandler {
	if cfg.MessagingPlatform == "watermill" {
		return NewWatermillMessageHandler(dcaSrv)
	}

	return NewSQSMessageHandler(dcaSrv, cfg)
}
