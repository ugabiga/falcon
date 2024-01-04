package repository

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type AuthenticationDynamoRepository struct {
	db *dynamodb.Client
}

func NewAuthenticationDynamoRepository(db *dynamodb.Client) *AuthenticationDynamoRepository {
	return &AuthenticationDynamoRepository{
		db: db,
	}
}
