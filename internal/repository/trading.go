package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/model"
	"time"
)

const (
	Separator             = "-"
	KeyPrefixTaskAccount  = "ta"
	KeyPrefixTask         = "task"
	KeyPrefixTaskHistory  = "th"
	EntityTypeTaskAccount = "trading_account"
	EntityTypeTask        = "task"
	EntityTypeTaskHistory = "task_history"
)

type TradingDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewTradingDynamoRepository(db *dynamodb.Client) *TradingDynamoRepository {
	return &TradingDynamoRepository{
		tableName: TradingTableName,
		db:        db,
	}
}

func (r TradingDynamoRepository) CreateTradingAccount(ctx context.Context, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.ID = r.encodeTradingAccountID(KeyPrefixTaskAccount, tradingAccount.UserID, tradingAccount.Exchange, tradingAccount.Key)
	tradingAccount.UpdatedAt = time.Now()
	tradingAccount.CreatedAt = time.Now()

	av, err := MarshalItem(tradingAccount)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: tradingAccount.UserID}
	av["sk"] = &types.AttributeValueMemberS{Value: tradingAccount.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskAccount}

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

func (r TradingDynamoRepository) UpdateTradingAccount(ctx context.Context, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.UpdatedAt = time.Now()

	av, err := MarshalItem(tradingAccount)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: tradingAccount.UserID}
	av["sk"] = &types.AttributeValueMemberS{Value: tradingAccount.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskAccount}

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

func (r TradingDynamoRepository) GetTradingAccount(ctx context.Context, userID, tradingAccountID string) (*model.TradingAccount, error) {
	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: userID},
				"sk": &types.AttributeValueMemberS{Value: tradingAccountID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	tradingAccount, err := UnmarshalItem[model.TradingAccount](result.Item)
	if err != nil {
		return nil, err
	}

	return tradingAccount, nil
}

func (r TradingDynamoRepository) GetTradingAccountsByUserID(ctx context.Context, userID string) ([]model.TradingAccount, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: userID},
				":sk": &types.AttributeValueMemberS{Value: KeyPrefixTaskAccount},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	var tradingAccounts []model.TradingAccount

	for _, item := range result.Items {
		tradingAccount, err := UnmarshalItem[model.TradingAccount](item)
		if err != nil {
			return nil, err
		}

		tradingAccounts = append(tradingAccounts, *tradingAccount)
	}

	return tradingAccounts, nil
}

func (r TradingDynamoRepository) CountTradingAccountsByUserID(ctx context.Context, userID string) (int, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: userID},
				":sk": &types.AttributeValueMemberS{Value: KeyPrefixTaskAccount},
			},
			Select: types.SelectCount,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(result.Count), nil
}

func (r TradingDynamoRepository) CreateTask(ctx context.Context, task model.Task) (*model.Task, error) {
	task.ID = r.encoding("task")
	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()

	av, err := MarshalItem(task)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: task.TradingAccountID}
	av["sk"] = &types.AttributeValueMemberS{Value: task.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTask}

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

func (r TradingDynamoRepository) UpdateTask(ctx context.Context, task model.Task) (*model.Task, error) {
	task.UpdatedAt = time.Now()

	av, err := MarshalItem(task)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: task.TradingAccountID}
	av["sk"] = &types.AttributeValueMemberS{Value: task.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTask}

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

func (r TradingDynamoRepository) GetTask(ctx context.Context, tradingAccountID, taskID string) (*model.Task, error) {
	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: tradingAccountID},
				"sk": &types.AttributeValueMemberS{Value: taskID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	task, err := UnmarshalItem[model.Task](result.Item)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r TradingDynamoRepository) GetTasksByTradingAccountID(ctx context.Context, tradingAccountID string) ([]model.Task, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: tradingAccountID},
				":sk": &types.AttributeValueMemberS{Value: KeyPrefixTask},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for _, item := range result.Items {
		task, err := UnmarshalItem[model.Task](item)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func (r TradingDynamoRepository) CountTasksByTradingID(ctx context.Context, tradingAccountID string) (int, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: tradingAccountID},
				":sk": &types.AttributeValueMemberS{Value: "task"},
			},
			Select: types.SelectCount,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(result.Count), nil
}

func (r TradingDynamoRepository) CreateTaskHistory(ctx context.Context, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.ID = r.encoding(KeyPrefixTaskHistory)
	taskHistory.UpdatedAt = time.Now()
	taskHistory.CreatedAt = time.Now()

	av, err := MarshalItem(taskHistory)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: taskHistory.TaskID}
	av["sk"] = &types.AttributeValueMemberS{Value: taskHistory.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskHistory}

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

	return &taskHistory, nil
}

func (r TradingDynamoRepository) UpdateTaskHistory(ctx context.Context, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.UpdatedAt = time.Now()

	av, err := MarshalItem(taskHistory)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: taskHistory.TaskID}
	av["sk"] = &types.AttributeValueMemberS{Value: taskHistory.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskHistory}

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

	return &taskHistory, nil
}

func (r TradingDynamoRepository) GetTaskHistory(ctx context.Context, taskID, taskHistoryID string) (*model.TaskHistory, error) {
	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: taskID},
				"sk": &types.AttributeValueMemberS{Value: taskHistoryID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	taskHistory, err := UnmarshalItem[model.TaskHistory](result.Item)
	if err != nil {
		return nil, err
	}

	return taskHistory, nil
}

func (r TradingDynamoRepository) GetTaskHistoriesByTaskID(ctx context.Context, taskID string) ([]model.TaskHistory, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: taskID},
				":sk": &types.AttributeValueMemberS{Value: KeyPrefixTaskHistory},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	var taskHistories []model.TaskHistory

	for _, item := range result.Items {
		taskHistory, err := UnmarshalItem[model.TaskHistory](item)
		if err != nil {
			return nil, err
		}

		taskHistories = append(taskHistories, *taskHistory)
	}

	return taskHistories, nil
}

func (r TradingDynamoRepository) encodeTradingAccountID(prefix, id, exchange, key string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", prefix, Separator, id, Separator, exchange, Separator, key)
}

func (r TradingDynamoRepository) encoding(prefix string) string {
	return prefix + Separator + uuid.New().String()
}
