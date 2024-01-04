package app

import (
	"github.com/ugabiga/falcon/internal/migration"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
)

type Tester struct {
	UserSrv           *service.UserService
	AuthenticationSrv *service.AuthenticationService
	TradingAccountSrv *service.TradingAccountService
	TaskSrv           *service.TaskService
	TaskHistorySrv    *service.TaskHistoryService
	DcaSrv            *service.DcaService
	Migration         *migration.Migration
}

func NewTester(
	userSrv *service.UserService,
	authenticationSrv *service.AuthenticationService,
	tradingAccountSrv *service.TradingAccountService,
	taskSrv *service.TaskService,
	taskHistorySrv *service.TaskHistoryService,
	dcaSrv *service.DcaService,
	mg *migration.Migration,
) Tester {
	return Tester{
		UserSrv:           userSrv,
		AuthenticationSrv: authenticationSrv,
		TradingAccountSrv: tradingAccountSrv,
		TaskSrv:           taskSrv,
		TaskHistorySrv:    taskHistorySrv,
		DcaSrv:            dcaSrv,
		Migration:         mg,
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
