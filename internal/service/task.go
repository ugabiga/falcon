package service

import (
	"context"
	"errors"
	"github.com/adhocore/gronx"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
	"github.com/ugabiga/falcon/internal/ent/tradingaccount"
	"github.com/ugabiga/falcon/internal/ent/user"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"time"

	"strconv"
	"strings"
)

const (
	TaskCreationLimit = 10
)

type TaskService struct {
	db *ent.Client
}

func NewTaskService(db *ent.Client) *TaskService {
	return &TaskService{db: db}
}

func (s TaskService) Create(ctx context.Context, userID int, input generated.CreateTaskInput) (*ent.Task, error) {
	if err := s.validateExceedLimit(ctx, userID); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	tradingAccount, err := s.db.TradingAccount.Query().
		Where(
			tradingaccount.UserID(userID),
			tradingaccount.ID(input.TradingAccountID),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	u, err := s.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := s.nextCronExecutionTime(cron, u.Timezone)
	if err != nil {
		return nil, err
	}

	return s.db.Task.Create().
		SetCurrency(input.Currency).
		SetSize(input.Size).
		SetSymbol(input.Symbol).
		SetCron(cron).
		SetNextExecutionTime(nextExecutionTime).
		SetType(input.Type).
		SetTradingAccountID(tradingAccount.ID).
		Save(ctx)
}

func (s TaskService) Update(ctx context.Context, userID int, taskID int, input generated.UpdateTaskInput) (*ent.Task, error) {
	if err := s.validateHours(input.Hours); err != nil {
		return nil, err
	}

	if err := s.validateUser(ctx, userID, taskID); err != nil {
		return nil, ErrDoNotHaveAccess
	}

	if err := s.validateCurrency(input.Currency); err != nil {
		return nil, err
	}

	u, err := s.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}

	cron := s.cronExpression(input.Hours, input.Days)
	nextExecutionTime, err := s.nextCronExecutionTime(cron, u.Timezone)
	if err != nil {
		return nil, err
	}

	return s.db.Task.UpdateOneID(taskID).
		SetCurrency(input.Currency).
		SetSize(input.Size).
		SetSymbol(input.Symbol).
		SetCron(cron).
		SetNextExecutionTime(nextExecutionTime).
		SetType(input.Type).
		SetIsActive(input.IsActive).
		SetParams(input.Params).
		Save(ctx)
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

func (s TaskService) GetByTradingAccount(ctx context.Context, tradingAccountID int) ([]*ent.Task, error) {
	return s.db.Task.Query().
		Where(
			task.TradingAccountID(tradingAccountID),
		).
		WithTradingAccount().
		All(ctx)
}

func (s TaskService) validateExceedLimit(ctx context.Context, userID int) error {
	count, err := s.db.TradingAccount.Query().
		Where(
			tradingaccount.UserID(userID),
		).
		Count(ctx)
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

func (s TaskService) nextCronExecutionTime(cron string, timezone string) (time.Time, error) {
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
