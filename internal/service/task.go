package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"strconv"
	"strings"
)

const (
	TaskCreationLimit = 10
)

type TaskService struct {
	repo *repository.DynamoRepository
}

func NewTaskService(
	repo *repository.DynamoRepository,
) *TaskService {
	return &TaskService{
		repo: repo,
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

	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
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
		return nil, ErrDoNotHaveAccess
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
	case "USD":
		return nil
	default:
		return ErrWrongCurrency
	}
}

func (s TaskService) cronExpression(hour string, days string) string {
	return "0 0 " + hour + " * * " + days
}
