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

	//err = m.createUserTable(ctx)
	//err = m.createAuthenticationTable(ctx)
	err = m.createTradingTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migration) DeleteAllTables(ctx context.Context) error {
	tables := []string{
		//repository.UserTableName,
		//repository.AuthenticationTableName,
		repository.TradingTableName,
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

func (m *Migration) createAuthenticationTable(ctx context.Context) error {
	tableName := repository.AuthenticationTableName
	_, err := m.dynamoDB.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("user_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("user-index"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("user_id"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
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

func (m *Migration) createUserTable(ctx context.Context) error {
	tableName := repository.UserTableName
	_, err := m.dynamoDB.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("name-index"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("name"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
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

func (m *Migration) createTradingTable(ctx context.Context) error {
	tableName := repository.TradingTableName
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
					ProjectionType: types.ProjectionTypeAll,
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
