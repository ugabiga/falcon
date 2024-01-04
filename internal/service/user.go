package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type UserService struct {
	db       *ent.Client
	userRepo *repository.UserDynamoRepository
}

func NewUserService(
	db *ent.Client,
	userRepo *repository.UserDynamoRepository,
) *UserService {
	return &UserService{
		db:       db,
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
func (s UserService) Update(ctx context.Context, userID string, user *model.User) (*model.User, error) {
	u, err := s.userRepo.Update(ctx, userID, user)
	if err != nil {
		return nil, err
	}

	return u, nil
}
