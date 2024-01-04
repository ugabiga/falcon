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

	err = m.createUserTable(ctx)
	err = m.createAuthenticationTable(ctx)
	err = m.createTradingAccountTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migration) DeleteAllTables(ctx context.Context) error {
	tables := []string{
		repository.AuthenticationTableName,
		repository.UserTableName,
		repository.TradingAccountTableName,
	}

	for _, table := range tables {
		if err := m.deleteTable(ctx, table); err != nil {
			log.Printf("error deleting table %s: %s", table, err)
			return err
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

func (m *Migration) createTradingAccountTable(ctx context.Context) error {
	tableName := repository.TradingAccountTableName
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
			//{
			//	AttributeName: aws.String("user_id"),
			//	KeyType:       types.KeyTypeRange,
			//},
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
		//LocalSecondaryIndexes: []types.LocalSecondaryIndex{
		//	{
		//		IndexName: aws.String("user-index"),
		//		KeySchema: []types.KeySchemaElement{
		//			{
		//				AttributeName: aws.String("id"),
		//				KeyType:       types.KeyTypeHash,
		//			},
		//			{
		//				AttributeName: aws.String("user_id"),
		//				KeyType:       types.KeyTypeRange,
		//			},
		//		},
		//		Projection: &types.Projection{
		//			ProjectionType: types.ProjectionTypeAll,
		//		},
		//	},
		//},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Printf("error creating table %s: %s", tableName, err)
		return err
	}
	return nil
}
