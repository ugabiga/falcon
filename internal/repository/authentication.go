package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ugabiga/falcon/internal/model"
)

const (
	AuthenticationTableName = "authentication"
)

type AuthenticationDynamoRepository struct {
	tableName string
	db        *dynamodb.Client
}

func NewAuthenticationDynamoRepository(db *dynamodb.Client) *AuthenticationDynamoRepository {
	return &AuthenticationDynamoRepository{
		tableName: AuthenticationTableName,
		db:        db,
	}
}

func (r AuthenticationDynamoRepository) Create(ctx context.Context, authentication model.Authentication) (*model.Authentication, error) {
	authentication.ID = r.encodeID(authentication.Provider, authentication.Identifier)

	av, err := MarshalItem(authentication)
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

	return &authentication, nil

}
func (r AuthenticationDynamoRepository) GetItem(ctx context.Context, provider, identifier string) (*model.Authentication, error) {
	raw, err := r.db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &r.tableName,
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: r.encodeID(provider, identifier)},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	result, err := UnmarshalItem[model.Authentication](raw.Item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r AuthenticationDynamoRepository) encodeID(provider, identifier string) string {
	return fmt.Sprintf("%s:%s", provider, identifier)
}
