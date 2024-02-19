package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ugabiga/falcon/internal/model"
	"math/rand"
	"strconv"
	"time"
)

var (
	ErrNotFound = fmt.Errorf("not_found")
)

const (
	TableName                 = "falcon"
	KeyPrefixUser             = "user"
	KeyPrefixAuthentication   = "auth"
	KeyPrefixTaskAccount      = "ta"
	KeyPrefixTask             = "task"
	KeyPrefixTaskHistory      = "th"
	KeyPrefixStaticIP         = "si"
	Separator                 = "-"
	GISNextExecutionTime      = "next_execution_time_GSI"
	IndexNextExecutionTimeKey = "next_execution_time"
	GISIndexIPAvailability    = "ip_availability_GSI"
	EntityTypeUser            = "user"
	EntityTypeAuthentication  = "authentication"
	EntityTypeTaskAccount     = "trading_account"
	EntityTypeTask            = "task"
	EntityTypeTaskHistory     = "task_history"
	EntityTypeStaticIP        = "static_ip"
)

type DynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewDynamoRepository(db *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{
		tableName: TableName,
		db:        db,
	}
}

func (r DynamoRepository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	user.ID = r.encoding(KeyPrefixUser)
	user.UpdatedAt = r.timeNow()
	user.CreatedAt = r.timeNow()

	av, err := MarshalItem(user)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: user.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: user.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeUser}

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

	return &user, nil
}

func (r DynamoRepository) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	user.UpdatedAt = r.timeNow()

	av, err := MarshalItem(user)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: user.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: user.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeUser}

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

	return &user, nil
}

func (r DynamoRepository) GetUser(ctx context.Context, userID string) (*model.User, error) {
	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: userID},
				"sk": &types.AttributeValueMemberS{Value: userID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, ErrNotFound
	}

	user, err := UnmarshalItem[model.User](result.Item)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r DynamoRepository) CreateAuthentication(ctx context.Context, authentication model.Authentication) (*model.Authentication, error) {
	authentication.ID = r.encodeAuthenticationID(KeyPrefixAuthentication, authentication.Provider, authentication.Identifier)
	authentication.CreatedAt = r.timeNow()
	authentication.UpdatedAt = r.timeNow()

	av, err := MarshalItem(authentication)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: authentication.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: authentication.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeAuthentication}

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

	return &authentication, nil
}

func (r DynamoRepository) UpdateAuthentication(ctx context.Context, authentication model.Authentication) (*model.Authentication, error) {
	authentication.UpdatedAt = r.timeNow()

	av, err := MarshalItem(authentication)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: authentication.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: authentication.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeAuthentication}

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

	return &authentication, nil
}

func (r DynamoRepository) GetAuthentication(ctx context.Context, provider, identifier string) (*model.Authentication, error) {
	authenticationID := r.encodeAuthenticationID(KeyPrefixAuthentication, provider, identifier)

	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: authenticationID},
				"sk": &types.AttributeValueMemberS{Value: authenticationID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, ErrNotFound
	}

	authentication, err := UnmarshalItem[model.Authentication](result.Item)
	if err != nil {
		return nil, err
	}

	return authentication, nil
}

func (r DynamoRepository) CreateTradingAccount(ctx context.Context, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.ID = r.EncodeTradingAccountID(tradingAccount.UserID, tradingAccount.Exchange, tradingAccount.Key)
	tradingAccount.UpdatedAt = r.timeNow()
	tradingAccount.CreatedAt = r.timeNow()

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

func (r DynamoRepository) UpdateTradingAccount(ctx context.Context, tradingAccount model.TradingAccount) (*model.TradingAccount, error) {
	tradingAccount.UpdatedAt = r.timeNow()

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

func (r DynamoRepository) GetTradingAccount(ctx context.Context, userID, tradingAccountID string) (*model.TradingAccount, error) {
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
		return nil, ErrNotFound
	}

	tradingAccount, err := UnmarshalItem[model.TradingAccount](result.Item)
	if err != nil {
		return nil, err
	}

	return tradingAccount, nil
}

func (r DynamoRepository) GetTradingAccountsByUserID(ctx context.Context, userID string) ([]model.TradingAccount, error) {
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

func (r DynamoRepository) CountTradingAccountsByUserID(ctx context.Context, userID string) (int, error) {
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

func (r DynamoRepository) DeleteTradingAccount(ctx context.Context, userID, tradingAccountID string) error {
	_, err := r.db.DeleteItem(
		ctx,
		&dynamodb.DeleteItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: userID},
				"sk": &types.AttributeValueMemberS{Value: tradingAccountID},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r DynamoRepository) CreateTask(ctx context.Context, task model.Task) (*model.Task, error) {
	task.ID = r.encoding(KeyPrefixTask)
	task.NextExecutionTime = task.NextExecutionTime.Truncate(time.Second)
	task.UpdatedAt = r.timeNow()
	task.CreatedAt = r.timeNow()

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

func (r DynamoRepository) UpdateTask(ctx context.Context, task model.Task) (*model.Task, error) {
	task.NextExecutionTime = task.NextExecutionTime.Truncate(time.Second)
	task.UpdatedAt = r.timeNow()

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

func (r DynamoRepository) GetTask(ctx context.Context, tradingAccountID, taskID string) (*model.Task, error) {
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
		return nil, ErrNotFound
	}

	task, err := UnmarshalItem[model.Task](result.Item)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r DynamoRepository) GetTasksByTradingAccountID(ctx context.Context, tradingAccountID string) ([]model.Task, error) {
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

func (r DynamoRepository) GetTasksByActiveNextExecutionTimeAndType(ctx context.Context, nextExecutionTime time.Time, taskType string) ([]model.Task, error) {
	formattedNextExecutionTime := nextExecutionTime.Format(time.RFC3339)

	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{GISNextExecutionTime}[0],
			KeyConditions: map[string]types.Condition{
				"next_execution_time": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: formattedNextExecutionTime},
					},
				},
			},
			//Is Active?
			FilterExpression: &[]string{"#v = :is_active AND #t = :task_type"}[0],
			ExpressionAttributeNames: map[string]string{
				"#v": "is_active",
				"#t": "type",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":is_active": &types.AttributeValueMemberBOOL{Value: true},
				":task_type": &types.AttributeValueMemberS{Value: taskType},
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

func (r DynamoRepository) GetTasksByActiveNextExecutionTime(ctx context.Context, nextExecutionTime time.Time) ([]model.Task, error) {
	formattedNextExecutionTime := nextExecutionTime.Format(time.RFC3339)

	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName: &r.tableName,
			IndexName: &[]string{GISNextExecutionTime}[0],
			KeyConditions: map[string]types.Condition{
				"next_execution_time": {
					ComparisonOperator: types.ComparisonOperatorEq,
					AttributeValueList: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: formattedNextExecutionTime},
					},
				},
			},
			//Is Active?
			FilterExpression: &[]string{"#v = :is_active"}[0],
			ExpressionAttributeNames: map[string]string{
				"#v": "is_active",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":is_active": &types.AttributeValueMemberBOOL{Value: true},
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

func (r DynamoRepository) CountTasksByTradingID(ctx context.Context, tradingAccountID string) (int, error) {
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

func (r DynamoRepository) DeleteTask(ctx context.Context, tradingAccountID, taskID string) error {
	_, err := r.db.DeleteItem(
		ctx,
		&dynamodb.DeleteItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: tradingAccountID},
				"sk": &types.AttributeValueMemberS{Value: taskID},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
func (r DynamoRepository) GetAllTaskHistories(ctx context.Context) ([]*model.TaskHistory, error) {
	result, err := r.db.Scan(
		ctx,
		&dynamodb.ScanInput{
			TableName:        &r.tableName,
			FilterExpression: &[]string{"entity_type = :entity_type"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":entity_type": &types.AttributeValueMemberS{Value: EntityTypeTaskHistory},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	var taskHistories []*model.TaskHistory

	for _, item := range result.Items {
		taskHistory, err := UnmarshalItem[model.TaskHistory](item)
		if err != nil {
			return nil, err
		}

		taskHistories = append(taskHistories, taskHistory)
	}

	return taskHistories, nil
}

func (r DynamoRepository) CreateTaskHistory(ctx context.Context, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.ID = r.encoding(KeyPrefixTaskHistory)
	taskHistory.UpdatedAt = r.timeNow()
	taskHistory.CreatedAt = r.timeNow()

	av, err := MarshalItem(taskHistory)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: taskHistory.TaskID}
	av["sk"] = &types.AttributeValueMemberS{Value: taskHistory.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskHistory}
	av["ttl"] = &types.AttributeValueMemberN{Value: strconv.FormatInt(time.Now().AddDate(0, 0, 30).Unix(), 10)}

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

func (r DynamoRepository) UpdateTaskHistory(ctx context.Context, taskHistory model.TaskHistory) (*model.TaskHistory, error) {
	taskHistory.UpdatedAt = r.timeNow()

	av, err := MarshalItem(taskHistory)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: taskHistory.TaskID}
	av["sk"] = &types.AttributeValueMemberS{Value: taskHistory.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeTaskHistory}
	av["ttl"] = &types.AttributeValueMemberN{Value: strconv.FormatInt(taskHistory.UpdatedAt.AddDate(0, 0, 30).Unix(), 10)}

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

func (r DynamoRepository) GetTaskHistory(ctx context.Context, taskID, taskHistoryID string) (*model.TaskHistory, error) {
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
		return nil, ErrNotFound
	}

	taskHistory, err := UnmarshalItem[model.TaskHistory](result.Item)
	if err != nil {
		return nil, err
	}

	return taskHistory, nil
}

func (r DynamoRepository) GetTaskHistoriesByTaskID(ctx context.Context, taskID string) ([]model.TaskHistory, error) {
	result, err := r.db.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:              &r.tableName,
			KeyConditionExpression: &[]string{"pk = :pk AND begins_with(sk, :sk)"}[0],
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: taskID},
				":sk": &types.AttributeValueMemberS{Value: KeyPrefixTaskHistory},
			},
			ScanIndexForward: &[]bool{false}[0],
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

func (r DynamoRepository) CreateStaticIP(ctx context.Context, staticIP model.StaticIP) (*model.StaticIP, error) {
	staticIP.ID = r.EncodeStaticIPID(staticIP.IPAddress)
	staticIP.UpdatedAt = r.timeNow()
	staticIP.CreatedAt = r.timeNow()

	if staticIP.ID == "" {
		return nil, ErrEmptyID
	}

	av, err := MarshalItem(staticIP)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: staticIP.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: staticIP.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeStaticIP}

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

	return &staticIP, nil
}

func (r DynamoRepository) UpdateStaticIP(ctx context.Context, staticIP model.StaticIP) (*model.StaticIP, error) {
	staticIP.UpdatedAt = r.timeNow()

	av, err := MarshalItem(staticIP)
	if err != nil {
		return nil, err
	}

	av["pk"] = &types.AttributeValueMemberS{Value: staticIP.ID}
	av["sk"] = &types.AttributeValueMemberS{Value: staticIP.ID}
	av["entity_type"] = &types.AttributeValueMemberS{Value: EntityTypeStaticIP}

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

	return &staticIP, nil
}

func (r DynamoRepository) GetStaticIP(ctx context.Context, staticIPID string) (*model.StaticIP, error) {
	result, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: staticIPID},
				"sk": &types.AttributeValueMemberS{Value: staticIPID},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, ErrNotFound
	}

	staticIP, err := UnmarshalItem[model.StaticIP](result.Item)
	if err != nil {
		return nil, err
	}

	return staticIP, nil
}

func (r DynamoRepository) CountUpStaticIPUsage(ctx context.Context, staticIPID string) error {
	_, err := r.db.UpdateItem(
		ctx,
		&dynamodb.UpdateItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: staticIPID},
				"sk": &types.AttributeValueMemberS{Value: staticIPID},
			},
			UpdateExpression: &[]string{"SET #v = #v + :val"}[0],
			ExpressionAttributeNames: map[string]string{
				"#v": "ip_usage_count",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":val": &types.AttributeValueMemberN{Value: "1"},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r DynamoRepository) CountDownStaticIPUsage(ctx context.Context, staticIPID string) error {
	_, err := r.db.UpdateItem(
		ctx,
		&dynamodb.UpdateItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: staticIPID},
				"sk": &types.AttributeValueMemberS{Value: staticIPID},
			},
			UpdateExpression: &[]string{"SET #v = #v - :val"}[0],
			ExpressionAttributeNames: map[string]string{
				"#v": "ip_usage_count",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":val": &types.AttributeValueMemberN{Value: "1"},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r DynamoRepository) GetStaticIPByAvailability(ctx context.Context) (*model.StaticIP, error) {
	result, err := r.db.Scan(
		ctx,
		&dynamodb.ScanInput{
			TableName:        &r.tableName,
			FilterExpression: &[]string{"#v = :is_available"}[0],
			ExpressionAttributeNames: map[string]string{
				"#v": "ip_availability",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":is_available": &types.AttributeValueMemberBOOL{Value: true},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, ErrNotFound
	}

	staticIP, err := UnmarshalItem[model.StaticIP](result.Items[0])
	if err != nil {
		return nil, err
	}

	return staticIP, nil

}

func (r DynamoRepository) encodeAuthenticationID(prefix, provider, identifier string) string {
	return fmt.Sprintf("%s%s%s%s%s", prefix, Separator, provider, Separator, identifier)
}

func (r DynamoRepository) EncodeTradingAccountID(userID, exchange, key string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", KeyPrefixTaskAccount, Separator, exchange, Separator, key, Separator, userID)
}

func (r DynamoRepository) EncodeStaticIPID(ip string) string {
	return fmt.Sprintf("%s%s%s", KeyPrefixStaticIP, Separator, ip)
}

func (r DynamoRepository) encoding(prefix string) string {
	return prefix + Separator + strconv.FormatInt(generateRowId(), 10)
}

func (r DynamoRepository) timeNow() time.Time {
	return time.Now().Truncate(time.Second)
}

func generateRowId() int64 {
	const customEpoch = 1300000000000
	shardId := rand.Intn(65) // Generates a random number between 0 and 64
	ts := time.Now().UnixNano()/int64(time.Millisecond) - customEpoch
	randID := rand.Intn(512)
	ts = ts << 6
	ts = ts + int64(shardId)
	return (ts * 512) + int64(randID)
}
