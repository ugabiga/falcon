package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/user"
)

type UserService struct {
	db *ent.Client
}

func NewUserService(
	db *ent.Client,
) *UserService {
	return &UserService{
		db: db,
	}
}

func (s UserService) GetByID(ctx context.Context, userID uint64) (*ent.User, error) {
	u, err := s.db.User.Query().
		Where(
			user.IDEQ(userID),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s UserService) Update(ctx context.Context, userID uint64, user *ent.User) (*ent.User, error) {
	u, err := s.db.User.UpdateOneID(userID).
		SetName(user.Name).
		SetTimezone(user.Timezone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
