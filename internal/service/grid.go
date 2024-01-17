package service

import (
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/repository"
)

type GridService struct {
	repo       *repository.DynamoRepository
	encryption *encryption.Encryption
}

func NewGridService(
	repo *repository.DynamoRepository,
	encryption *encryption.Encryption,
) *GridService {
	return &GridService{
		repo:       repo,
		encryption: encryption,
	}
}

func (s GridService) GetTarget() ([]TaskOrderInfo, error) {
	return nil, nil
}
