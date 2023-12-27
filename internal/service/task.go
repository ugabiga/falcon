package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
	"github.com/ugabiga/falcon/internal/ent/tradingaccount"
	"strconv"
	"strings"
)

type TaskService struct {
	db *ent.Client
}

func NewTaskService(db *ent.Client) *TaskService {
	return &TaskService{db: db}
}

func (s TaskService) GetByTradingAccount(ctx context.Context, tradingAccountID int) ([]*ent.Task, error) {
	return s.db.Task.Query().
		Where(
			task.TradingAccountID(tradingAccountID),
		).
		WithTradingAccount().
		All(ctx)
}

func (s TaskService) Create(ctx context.Context, userID int, tradingAccountID int, hours string, typeArg string) (*ent.Task, error) {
	tradingAccount, err := s.db.TradingAccount.Query().
		Where(
			tradingaccount.UserID(userID),
			tradingaccount.ID(tradingAccountID),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	if err = s.validateHours(hours); err != nil {
		return nil, err
	}
	cron := "0 0 " + hours + " * * *"

	return s.db.Task.Create().
		SetCron(cron).
		SetType(typeArg).
		SetTradingAccountID(tradingAccount.ID).
		Save(ctx)
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

func (s TaskService) Update(ctx context.Context, userID int, id int, hours string, typeArg string, isActive bool) (*ent.Task, error) {
	if err := s.validateHours(hours); err != nil {
		return nil, err
	}

	cron := "0 0 " + hours + " * * *"

	return s.db.Task.UpdateOneID(id).
		SetCron(cron).
		SetType(typeArg).
		SetIsActive(isActive).
		Save(ctx)
}
