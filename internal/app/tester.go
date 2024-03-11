package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/messaging"
	"github.com/ugabiga/falcon/internal/migration"
	"github.com/ugabiga/falcon/internal/model"
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
	MessageHandler    messaging.MessageHandler
	MigrationSrv      *service.MigrationService
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
	messageHandler messaging.MessageHandler,
	migrationSrv *service.MigrationService,
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
		MessageHandler:    messageHandler,
		MigrationSrv:      migrationSrv,
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

func (t Tester) CreateStaticIP(ctx context.Context, tt *testing.T) {
	ipAddress := "127.0.0.1"

	_, err := t.Repository.CreateStaticIP(ctx, model.StaticIP{
		IPAddress:      ipAddress,
		IPAvailability: true,
		IPUsageCount:   0,
	})

	if err != nil {
		tt.Fatal(err)
	}
}

func (t Tester) CreateOrGetTestUser(ctx context.Context, tt *testing.T) *model.User {
	authentication, user, err := t.AuthenticationSrv.SignUp(
		ctx,
		"google",
		uuid.New().String(),
		uuid.New().String(),
		"new_user",
	)
	if err != nil {
		tt.Fatal(err)
	}

	if authentication == nil {
		tt.Fatal("authentication is nil")
	}

	if user == nil {
		tt.Fatal("user is nil")
	}

	return user
}

func (t Tester) CreateTestTradingAccount(
	ctx context.Context,
	tt *testing.T,
	userID,
	exchange,
	key,
	secret string,
) *model.TradingAccount {
	t.CreateStaticIP(ctx, tt)

	tradingAccount, err := t.TradingAccountSrv.Create(
		ctx,
		userID,
		"test",
		exchange,
		key,
		secret,
		"",
	)
	if err != nil {
		tt.Fatal(err)
	}

	return tradingAccount
}

func (t Tester) CreateTestTasks(
	ctx context.Context,
	tt *testing.T,
	tradingAccount *model.TradingAccount,
	userID string,
) []model.Task {

	//Create DCA Task
	dcaTask, err := t.TaskSrv.Create(
		ctx,
		userID,
		generated.CreateTaskInput{
			TradingAccountID: tradingAccount.ID,
			Currency:         "KRW",
			Size:             0.001,
			Symbol:           "BTC",
			Days:             "1,2,3,4,5,6,7",
			Hours:            "18",
			Type:             model.TaskTypeDCA,
			Params:           nil,
		},
	)
	if err != nil {
		tt.Fatal(err)
	}

	//Create Grid Task
	gridTask, err := t.TaskSrv.Create(
		ctx,
		userID,
		generated.CreateTaskInput{
			TradingAccountID: tradingAccount.ID,
			Currency:         "KRW",
			Size:             0.001,
			Symbol:           "BTC",
			Days:             "1,2,3,4,5,6,7",
			Hours:            "18",
			Type:             model.TaskTypeLongGrid,
			Params: model.TaskGridParams{
				GapPercent: 2,
				Quantity:   2,
			}.ToParams(),
		},
	)
	if err != nil {
		tt.Fatal(err)
	}

	return []model.Task{*dcaTask, *gridTask}
}
