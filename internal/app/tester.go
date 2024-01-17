package app

import (
	"context"
	"github.com/google/uuid"
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
	tt *testing.T, userID,
	exchange string,
	key string,
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
