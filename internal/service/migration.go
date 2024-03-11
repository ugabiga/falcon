package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
)

type MigrationService struct {
	repo       *repository.DynamoRepository
	encryption *encryption.Encryption
}

func NewMigrationService(
	repo *repository.DynamoRepository,
	encryption *encryption.Encryption,
) *MigrationService {
	return &MigrationService{
		repo:       repo,
		encryption: encryption,
	}
}
func (s MigrationService) Migrate() error {
	ctx := context.Background()
	log.Printf("Migrating grid params")
	if err := s.MigrateGridParams(ctx); err != nil {
		return err
	}

	return nil
}

func (s MigrationService) MigrateGridParams(ctx context.Context) error {
	tasks, err := s.repo.ScanTasksByType(ctx, model.TaskTypeLongGrid)
	if err != nil {
		return err
	}

	log.Printf("Found %d tasks", len(tasks))
	for _, task := range tasks {
		if task.Type != model.TaskTypeLongGrid {
			continue
		}

		gridParams, err := task.GridParams()
		if err != nil {
			return err
		}

		newParams := model.TaskGridParamsV2{
			GapPercent:           gridParams.GapPercent,
			Quantity:             gridParams.Quantity,
			UseIncrementalSize:   false,
			IncrementalSize:      0,
			DeletePreviousOrders: true,
		}

		task.Params = newParams.ToParams()
		_, err = s.repo.UpdateTask(ctx, task)
		if err != nil {
			return err
		}
		log.Printf("Updated task %s", task.ID)
	}

	return nil
}
