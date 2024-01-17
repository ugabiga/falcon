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
	gridSrv *service.GridService,
) MessageHandler {
	return NewSQSMessageHandler(cfg, dcaSrv, gridSrv)
}
