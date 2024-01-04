package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/model"
	"time"
)

const (
	UserTableName = "user"
)

type UserDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewUserDynamoRepository(db *dynamodb.Client) *UserDynamoRepository {
	return &UserDynamoRepository{
		tableName: UserTableName,
		db:        db,
	}
}

func (r UserDynamoRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	user.ID = r.encodeID()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()

	av, err := MarshalItem(user)
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

	return &user, nil
}

func (r UserDynamoRepository) Get(ctx context.Context, id string) (*model.User, error) {
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

	user, err := UnmarshalItem[model.User](av.Item)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserDynamoRepository) encodeID() string {
	return uuid.New().String()
}
