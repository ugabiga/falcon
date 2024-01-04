package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ugabiga/falcon/internal/model"
	"time"
)

type TradingAccountDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewTradingAccountDynamoRepository(db *dynamodb.Client) *TradingAccountDynamoRepository {
	return &TradingAccountDynamoRepository{
		tableName: TradingAccountTableName,
		db:        db,
	}
}

func (r TradingAccountDynamoRepository) encodeID(userID, exchange, key string) string {
	return fmt.Sprintf("%s:%s:%s", userID, exchange, key)
}

func (r TradingAccountDynamoRepository) Create(ctx context.Context, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.ID = r.encodeID(tradingAccount.UserID, tradingAccount.Exchange, tradingAccount.Key)
	tradingAccount.UpdatedAt = time.Now()
	tradingAccount.CreatedAt = time.Now()

	av, err := MarshalItem(tradingAccount)
	if err != nil {
		return nil, err
	}

	_, err = r.db.PutItem(
		ctx,
		&dynamodb.PutItemInput{
			TableName: &r.tableName,
			Item:      av,
		},
	)
	if err != nil {
		return nil, err
	}

	return &tradingAccount, nil
}

func (r TradingAccountDynamoRepository) Update(ctx context.Context, tradingAccountID string, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.ID = tradingAccountID
	tradingAccount.UpdatedAt = time.Now()

	av, err := MarshalItem(tradingAccount)
	if err != nil {
		return nil, err
	}

	_, err = r.db.PutItem(
		ctx,
		&dynamodb.PutItemInput{
			TableName: &r.tableName,
			Item:      av,
		},
	)
	if err != nil {
		return nil, err
	}

	return &tradingAccount, nil
}

func (r TradingAccountDynamoRepository) Count(ctx context.Context, userID string) (int, error) {
	out, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{"user-index"}[0],
			KeyConditions: map[string]types.Condition{
				"user_id": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: userID},
					},
				},
			},
		},
	)
	if err != nil {
		return 0, err
	}

	return int(out.Count), nil
}

func (r TradingAccountDynamoRepository) GetByUserID(ctx context.Context, userID string) ([]model.TradingAccount, error) {
	out, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{"user-index"}[0],
			KeyConditions: map[string]types.Condition{
				"user_id": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: userID},
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	items, err := UnmarshalItems[model.TradingAccount](out.Items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r TradingAccountDynamoRepository) GetByID(ctx context.Context, tradingAccountID string) (*model.TradingAccount, error) {
	out, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: tradingAccountID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	item, err := UnmarshalItem[model.TradingAccount](out.Item)
	if err != nil {
		return nil, err
	}

	return item, nil

}
