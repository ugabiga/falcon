package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/handler/v1/request"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
)

type UserService struct {
	repo *repository.DynamoRepository
}

func NewUserService(
	tradingRepo *repository.DynamoRepository,
) *UserService {
	return &UserService{
		repo: tradingRepo,
	}
}

type UserWithOptions struct {
	WithAuthentications bool
	WithTradingAccounts bool
}

func (s UserService) GetByID(ctx context.Context, userID string) (*model.User, error) {
	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s UserService) Update(ctx context.Context, userID string, userInput generated.UpdateUserInput) (*model.User, error) {
	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Name = userInput.Name
	u.Timezone = userInput.Timezone

	updateUser, err := s.repo.UpdateUser(ctx, *u)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}
func (s UserService) UpdateV1(ctx context.Context, userID string, userRequest request.UpdateUserRequest) (*model.User, error) {
	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Name = userRequest.Name
	u.Timezone = userRequest.Timezone

	updateUser, err := s.repo.UpdateUser(ctx, *u)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}
