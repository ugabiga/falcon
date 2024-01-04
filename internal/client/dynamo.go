package client

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/ugabiga/falcon/pkg/config"
	"log"
)

func NewDynamoClient(c *cfg.Config) (*dynamodb.Client, error) {
	ctx := context.Background()
	dynamoConfig, err := newDynamoConfig(ctx, c.AWSRegion, c.DynamoIsLocal)
	if err != nil {
		return nil, err
	}

	db := dynamodb.NewFromConfig(dynamoConfig)

	return db, nil
}

func newDynamoConfig(ctx context.Context, region string, isLocal bool) (aws.Config, error) {
	var dynamoClientConfig aws.Config
	var err error

	log.Println("isLocal: ", isLocal)

	switch isLocal {
	case true:
		//Local
		accessKey := "accessKey"
		secretAccessKey := "secretAccessKey"
		sessionToken := "sessionToken"
		endpoint := "http://localhost:8000"

		dynamoClientConfig, err = config.LoadDefaultConfig(
			ctx,
			config.WithRegion(region),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: endpoint}, nil
				},
			)),
			config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     accessKey,
					SecretAccessKey: secretAccessKey,
					SessionToken:    sessionToken,
					Source:          "Hard-coded credentials; values are irrelevant for local DynamoDB",
				},
			}),
		)
		if err != nil {
			return aws.Config{}, err
		}

	case false:
		//Remote AWS
		dynamoClientConfig, err = config.LoadDefaultConfig(
			ctx,
			config.WithRegion(region),
		)
		if err != nil {
			return aws.Config{}, err
		}

	default:
		log.Printf("invalid value for isLocal: %v", isLocal)
		return aws.Config{}, errors.New("invalid value for isLocal")
	}

	return dynamoClientConfig, nil
}
