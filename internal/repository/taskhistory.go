package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/model"
	"time"
)

type TaskHistoryDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewTaskHistoryDynamoRepository(db *dynamodb.Client) *TaskHistoryDynamoRepository {
	return &TaskHistoryDynamoRepository{
		tableName: TaskHistoryTableName,
		db:        db,
	}
}

func (r TaskHistoryDynamoRepository) encodeID() string {
	return uuid.New().String()
}

func (r TaskHistoryDynamoRepository) Create(ctx context.Context, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.ID = r.encodeID()
	taskHistory.UpdatedAt = time.Now()
	taskHistory.CreatedAt = time.Now()

	av, err := MarshalItem(taskHistory)
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

	return &taskHistory, nil
}

func (r TaskHistoryDynamoRepository) Update(ctx context.Context, taskHistoryID string, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.ID = taskHistoryID
	taskHistory.UpdatedAt = time.Now()

	av, err := MarshalItem(taskHistory)
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

	return &taskHistory, nil
}

func (r TaskHistoryDynamoRepository) Get(ctx context.Context, id string) (*model.TaskHistory, error) {
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

	taskHistory, err := UnmarshalItem[model.TaskHistory](av.Item)
	if err != nil {
		return nil, err
	}

	return taskHistory, nil
}

func (r TaskHistoryDynamoRepository) GetByTaskID(ctx context.Context, taskID string) ([]model.TaskHistory, error) {
	out, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{"task-index"}[0],
			KeyConditions: map[string]types.Condition{
				"task_id": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: taskID},
					},
				},
			},
			Limit: &[]int32{10}[0],
		},
	)
	if err != nil {
		return nil, err
	}

	taskHistories := make([]model.TaskHistory, len(out.Items))
	for i, item := range out.Items {
		taskHistory, err := UnmarshalItem[model.TaskHistory](item)
		if err != nil {
			return nil, err
		}

		taskHistories[i] = *taskHistory
	}

	return taskHistories, nil
}
