package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	TaskCreationLimit = 10
)

type TaskService struct {
	repo       *repository.DynamoRepository
	encryption *encryption.Encryption
}

func NewTaskService(
	repo *repository.DynamoRepository,
	encryption *encryption.Encryption,
) *TaskService {
	return &TaskService{
		repo:       repo,
		encryption: encryption,
	}
}

func (s TaskService) Create(ctx context.Context, userID string, input generated.CreateTaskInput) (*model.Task, error) {
	if err := s.validateExceedLimit(ctx, input.TradingAccountID); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	if err := s.validateSize(ctx, userID, input.TradingAccountID, input.Currency, input.Symbol, input.Size); err != nil {
		return nil, err
	}

	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := nextCronExecutionTime(cron, u.Timezone)
	if err != nil {
		return nil, err
	}

	newTask := model.Task{
		UserID:            userID,
		TradingAccountID:  input.TradingAccountID,
		Currency:          input.Currency,
		Size:              input.Size,
		Symbol:            input.Symbol,
		Cron:              cron,
		NextExecutionTime: nextExecutionTime,
		IsActive:          true,
		Type:              input.Type,
		Params:            input.Params,
	}

	t, err := s.repo.CreateTask(ctx, newTask)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s TaskService) Update(ctx context.Context, userID string, tradingAccountID string, taskID string, input generated.UpdateTaskInput) (*model.Task, error) {
	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := nextCronExecutionTime(cron, u.Timezone)
	if err != nil {
		return nil, err
	}

	//nextExecutionTime = time.Now().UTC().Add(time.Minute * 1).Truncate(time.Minute)

	t, err := s.repo.GetTask(ctx, tradingAccountID, taskID)
	if err != nil {
		return nil, err
	}

	if t.UserID != userID {
		return nil, ErrUnAuthorizedAction
	}

	updateTask := model.Task{
		ID:                t.ID,
		UserID:            t.UserID,
		TradingAccountID:  t.TradingAccountID,
		Currency:          input.Currency,
		Size:              input.Size,
		Symbol:            input.Symbol,
		Cron:              cron,
		NextExecutionTime: nextExecutionTime,
		IsActive:          input.IsActive,
		Type:              input.Type,
		Params:            input.Params,
		CreatedAt:         t.CreatedAt,
	}

	return s.repo.UpdateTask(ctx, updateTask)
}

func (s TaskService) Delete(ctx context.Context, userID string, tradingAccountID string, taskID string) error {
	t, err := s.repo.GetTask(ctx, tradingAccountID, taskID)
	if err != nil {
		return err
	}

	if t.UserID != userID {
		return ErrUnAuthorizedAction
	}

	return s.repo.DeleteTask(ctx, tradingAccountID, taskID)
}

func (s TaskService) GetByTradingAccount(ctx context.Context, tradingAccountID string) ([]model.Task, error) {
	tasks, err := s.repo.GetTasksByTradingAccountID(ctx, tradingAccountID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s TaskService) MigrateGridParams(ctx context.Context) error {
	tasks, err := s.repo.ScanTasksByType(ctx, model.TaskTypeLongGrid)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Type != model.TaskTypeLongGrid {
			continue
		}

		gridParams, err := task.GridParams()
		if err != nil {
			return err
		}

		newParams := model.TaskGridParamsV2{
			GapPercent:                 gridParams.GapPercent,
			Quantity:                   gridParams.Quantity,
			ShouldDeletePreviousOrders: true,
		}

		task.Params = newParams.ToParams()
		_, err = s.repo.UpdateTask(ctx, task)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TaskService) validateExceedLimit(ctx context.Context, tradingAccountID string) error {
	count, err := s.repo.CountTasksByTradingID(ctx, tradingAccountID)
	if err != nil {
		return err
	}

	if count >= TaskCreationLimit {
		return ErrExceedLimit
	}

	return nil
}

func (s TaskService) validateHours(hours string) error {
	splitHours := strings.Split(hours, ",")
	for _, hour := range splitHours {
		intHour, err := strconv.Atoi(hour)
		if err != nil {
			return errors.New("hours should be integers")
		}
		if intHour < 0 || intHour > 23 {
			return errors.New("hours should be in the range 0-23")
		}
	}
	return nil
}

func (s TaskService) validateCurrency(currency string) error {
	// currency code ISO 4217
	switch currency {
	case "KRW":
		return nil
	case "USDT":
		return nil
	default:
		return ErrWrongCurrency
	}
}

func (s TaskService) validateSize(ctx context.Context, userID, tradingAccountID, currency, symbol string, size float64) error {
	if size < 0 {
		return ErrSizeNotSatisfiedMinimumSize
	}

	tradingAccount, err := s.repo.GetTradingAccount(ctx, userID, tradingAccountID)
	if err != nil {
		return err
	}

	switch tradingAccount.Exchange {
	case model.ExchangeUpbit:
		return s.validateUpbitSize(ctx, currency, symbol, size)
	case model.ExchangeBinanceSpot:
		return s.validateBinanceSize(ctx, tradingAccount, currency, symbol, size)
	case model.ExchangeBinanceFutures:
		return s.validateBinanceSize(ctx, tradingAccount, currency, symbol, size)
	default:
		return ErrWrongExchange
	}
}

func (s TaskService) cronExpression(hour string, days string) string {
	return "0 0 " + hour + " * * " + days
}

func (s TaskService) validateUpbitSize(ctx context.Context, currency, symbol string, size float64) error {
	upbitClient := client.NewUpbitClient("", "")
	minimumUpbitCost := 5000
	upbitSymbol := currency + "-" + symbol
	ticker, err := upbitClient.TickerPublic(ctx, upbitSymbol)
	if err != nil {
		return err
	}

	minimumSize := float64(minimumUpbitCost) / ticker.TradePrice
	if size < minimumSize {
		return errors.New(ErrSizeNotSatisfiedMinimumSize.Error() + fmt.Sprintf("#%f", minimumSize))
	}

	return nil
}

func (s TaskService) validateBinanceSize(ctx context.Context, tradingAccount *model.TradingAccount, currency, symbol string, size float64) error {
	return nil
	// There is ip restriction for binance futures
	// So, we can't use binance futures api with the current setup

	secret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}
	binanceClient := client.NewBinanceFutureClient(tradingAccount.Key, secret, false)
	binanceSymbol := symbol + currency
	minimumBinanceCost := 5.0

	ticker, err := binanceClient.Ticker(ctx, binanceSymbol)
	if err != nil {
		return err
	}
	tickerPrice := str.New(ticker.Price).ToFloat64Default(0)

	position, err := binanceClient.PositionWithoutSideIncludeEmpty(ctx, binanceSymbol)
	if err != nil {
		return err
	}
	leverage := str.New(position.Leverage).ToIntDefault(0)

	_, stepSizeStr, err := binanceClient.TickAndStepSize(ctx, binanceSymbol)
	if err != nil {
		return err
	}
	stepSizeDecimalCount := str.New(stepSizeStr).CountDecimalCount()

	minimumCost := math.Round(minimumBinanceCost * float64(leverage))
	minimumSize := math.Round(minimumCost/tickerPrice*math.Pow10(stepSizeDecimalCount)) / math.Pow10(stepSizeDecimalCount)

	if size < minimumSize {
		return errors.New(ErrSizeNotSatisfiedMinimumSize.Error() +
			"#" +
			strconv.FormatFloat(minimumSize, 'f', stepSizeDecimalCount, 64),
		)
	}

	return nil
}
