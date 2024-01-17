package messaging

import (
	"github.com/ugabiga/falcon/internal/service"
)

type TaskOrderInfoMessage struct {
	TaskOrderInfo service.TaskOrderInfo
}
