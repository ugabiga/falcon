package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/ent"
	"log"
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

func (s UserService) GetUser(ctx context.Context) error {
	u, err := s.db.User.Query().First(ctx)
	if err != nil {
		return err
	}

	log.Printf("User: %v", u.Name)

	return nil
}
