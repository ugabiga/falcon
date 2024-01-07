package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type UserService struct {
	tradingRepo *repository.TradingDynamoRepository
}

func NewUserService(
	tradingRepo *repository.TradingDynamoRepository,
) *UserService {
	return &UserService{
		tradingRepo: tradingRepo,
	}
}

type UserWithOptions struct {
	WithAuthentications bool
	WithTradingAccounts bool
}

func (s UserService) GetByID(ctx context.Context, userID string) (*model.User, error) {
	u, err := s.tradingRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s UserService) Update(ctx context.Context, userID string, userInput generated.UpdateUserInput) (*model.User, error) {
	u, err := s.tradingRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Name = userInput.Name
	u.Timezone = userInput.Timezone

	updateUser, err := s.tradingRepo.UpdateUser(ctx, *u)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}
