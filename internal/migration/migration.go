package migration

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
)

type Migration struct {
	dynamoDB *dynamodb.Client
}

func NewMigration(dynamoDB *dynamodb.Client) *Migration {
	return &Migration{
		dynamoDB: dynamoDB,
	}
}

func (m *Migration) Migrate(afterDelete bool) error {
	var err error

	ctx := context.Background()

	if afterDelete {
		err = m.DeleteAllTables(ctx)
	}

	err = m.createTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migration) DeleteAllTables(ctx context.Context) error {
	tables := []string{
		repository.TableName,
	}

	for _, table := range tables {
		if err := m.deleteTable(ctx, table); err != nil {
			log.Printf("error deleting table %s: %s", table, err)
			continue
		}
	}

	return nil
}

func (m *Migration) deleteTable(ctx context.Context, tableName string) error {
	_, err := m.dynamoDB.DeleteTable(ctx, &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Migration) createTable(ctx context.Context) error {
	tableName := repository.TableName
	_, err := m.dynamoDB.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String(repository.IndexNextExecutionTimeKey),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String(repository.GISNextExecutionTime),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String(repository.IndexNextExecutionTimeKey),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeInclude,
					NonKeyAttributes: []string{
						"pk",
						"sk",
						"id",
						"user_id",
						"trading_account_id",
						"currency",
						"size",
						"symbol",
						"cron",
						"next_execution_time",
						"is_active",
						"type",
						"params",
					},
				},
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Printf("error creating table %s: %s", tableName, err)
		return err
	}
	return nil
}
