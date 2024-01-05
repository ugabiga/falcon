package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type UserService struct {
	userRepo *repository.UserDynamoRepository
}

func NewUserService(
	userRepo *repository.UserDynamoRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type UserWithOptions struct {
	WithAuthentications bool
	WithTradingAccounts bool
}

func (s UserService) GetByID(ctx context.Context, userID string) (*model.User, error) {
	u, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s UserService) Update(ctx context.Context, userID string, userInput generated.UpdateUserInput) (*model.User, error) {
	u, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Name = userInput.Name
	u.Timezone = userInput.Timezone

	updateUser, err := s.userRepo.Update(ctx, userID, u)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}
