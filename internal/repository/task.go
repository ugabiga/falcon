package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/model"
	"time"
)

type TaskDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewTaskDynamoRepository(db *dynamodb.Client) *TaskDynamoRepository {
	return &TaskDynamoRepository{
		tableName: TaskTableName,
		db:        db,
	}
}

func (r TaskDynamoRepository) Create(ctx context.Context, task model.Task) (*model.Task, error) {
	task.ID = r.encodeID()
	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()

	av, err := MarshalItem(task)
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

	return &task, nil
}

func (r TaskDynamoRepository) Update(ctx context.Context, taskID string, task model.Task) (*model.Task, error) {
	task.ID = taskID
	task.UpdatedAt = time.Now()

	av, err := MarshalItem(task)
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

	return &task, nil
}

func (r TaskDynamoRepository) Get(ctx context.Context, id string) (*model.Task, error) {
	av, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: id},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	task, err := UnmarshalItem[model.Task](av.Item)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r TaskDynamoRepository) GetByTradingAccountID(ctx context.Context, tradingAccountID string) ([]model.Task, error) {
	out, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{"trading-account-index"}[0],
			KeyConditions: map[string]types.Condition{
				"trading_account_id": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: tradingAccountID},
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	items, err := UnmarshalItems[model.Task](out.Items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r TaskDynamoRepository) CountByUserID(ctx context.Context, userID string) (int, error) {
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

func (r TaskDynamoRepository) encodeID() string {
	return uuid.New().String()
}
