package app

import (
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/migration"
	"github.com/ugabiga/falcon/internal/repository"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"testing"
)

type Tester struct {
	Cfg               *config.Config
	UserSrv           *service.UserService
	AuthenticationSrv *service.AuthenticationService
	TradingAccountSrv *service.TradingAccountService
	TaskSrv           *service.TaskService
	TaskHistorySrv    *service.TaskHistoryService
	DcaSrv            *service.DcaService
	Migration         *migration.Migration
	Repository        *repository.DynamoRepository
	GridSrv           *service.GridService
	SQSHandler        *messaging.SQSMessageHandler
}

func NewTester(
	cfg *config.Config,
	userSrv *service.UserService,
	authenticationSrv *service.AuthenticationService,
	tradingAccountSrv *service.TradingAccountService,
	taskSrv *service.TaskService,
	taskHistorySrv *service.TaskHistoryService,
	dcaSrv *service.DcaService,
	mg *migration.Migration,
	tradingRepository *repository.DynamoRepository,
	gridSrv *service.GridService,
	sqsHandler *messaging.SQSMessageHandler,
) Tester {
	return Tester{
		Cfg:               cfg,
		UserSrv:           userSrv,
		AuthenticationSrv: authenticationSrv,
		TradingAccountSrv: tradingAccountSrv,
		TaskSrv:           taskSrv,
		TaskHistorySrv:    taskHistorySrv,
		DcaSrv:            dcaSrv,
		Migration:         mg,
		Repository:        tradingRepository,
		GridSrv:           gridSrv,
		SQSHandler:        sqsHandler,
	}
}

func (t Tester) ResetTables(tt *testing.T) {
	if err := t.Migration.Migrate(true); err != nil {
		tt.Fatal(err)
	}
}
func (t Tester) CleanUp(tt *testing.T) {
	if err := t.Migration.Migrate(true); err != nil {
		tt.Fatal(err)
	}
}
func (t Tester) CreateOrGetTestUser(tt *testing.T) {

}
