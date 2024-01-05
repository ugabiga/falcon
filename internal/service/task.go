package service

import (
	"context"
	"errors"
	"github.com/adhocore/gronx"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"time"

	"strconv"
	"strings"
)

const (
	TaskCreationLimit = 10
)

type TaskService struct {
	db       *ent.Client
	taskRepo *repository.TaskDynamoRepository
	userRepo *repository.UserDynamoRepository
}

func NewTaskService(
	db *ent.Client,
	taskRepo *repository.TaskDynamoRepository,
	userRepo *repository.UserDynamoRepository,
) *TaskService {
	return &TaskService{
		db:       db,
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s TaskService) Create(ctx context.Context, userID string, input generated.CreateTaskInput) (*model.Task, error) {
	if err := s.validateExceedLimit(ctx, userID); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	u, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := NextCronExecutionTime(cron, u.Timezone)
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

	t, err := s.taskRepo.Create(ctx, newTask)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s TaskService) Update(ctx context.Context, userID string, taskID string, input generated.UpdateTaskInput) (*model.Task, error) {
	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	u, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := NextCronExecutionTime(cron, u.Timezone)
	if err != nil {
		return nil, err
	}

	t, err := s.taskRepo.Get(ctx, taskID)
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

	return s.taskRepo.Update(ctx, taskID, updateTask)
}

func (s TaskService) GetWithTaskHistory(ctx context.Context, userID, Id int) (*ent.Task, error) {
	if err := s.validateUser(ctx, userID, Id); err != nil {
		return nil, err
	}

	return s.db.Task.Query().
		Where(
			task.ID(Id),
		).
		WithTaskHistories().
		First(ctx)
}

func (s TaskService) GetByTradingAccount(ctx context.Context, tradingAccountID string) ([]model.Task, error) {
	tasks, err := s.taskRepo.GetByTradingAccountID(ctx, tradingAccountID)
	if err != nil {
		return nil, err
	}

	return tasks, nil

	//return s.db.Task.Query().
	//	Where(
	//		task.TradingAccountID(tradingAccountID),
	//	).
	//	WithTradingAccount().
	//	All(ctx)
}

func (s TaskService) Get(ctx context.Context, userID string, taskID string) (*model.Task, error) {
	t, err := s.taskRepo.Get(ctx, taskID)
	if err != nil {
		return nil, err
	}

	if t.UserID != userID {
		return nil, ErrDoNotHaveAccess
	}

	return t, nil
}

func (s TaskService) validateExceedLimit(ctx context.Context, userID string) error {
	count, err := s.taskRepo.CountByUserID(ctx, userID)
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

func (s TaskService) validateUser(ctx context.Context, userID int, id int) error {
	targetTask, err := s.db.Task.Query().Where(
		task.ID(id),
	).WithTradingAccount().First(ctx)
	if err != nil {
		return err
	}

	if targetTask.Edges.TradingAccount.UserID != userID {
		return ErrDoNotHaveAccess
	}

	return err
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

func NextCronExecutionTime(cron string, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	localTime := time.Now().In(location)
	nextTime, err := gronx.NextTickAfter(cron, localTime, true)
	if err != nil {
		return time.Time{}, err
	}

	return nextTime.UTC(), nil
}
