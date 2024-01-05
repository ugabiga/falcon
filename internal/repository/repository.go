package repository

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	UserTableName           = "user"
	AuthenticationTableName = "authentication"
	TradingAccountTableName = "trading_account"
	TaskTableName           = "task"
	TaskHistoryTableName    = "task_history"
)

func UnmarshalItem[T any](item map[string]types.AttributeValue) (*T, error) {
	result := new(T)

	if err := attributevalue.UnmarshalMapWithOptions(item, result,
		func(options *attributevalue.DecoderOptions) {
			options.TagKey = "json"
		}); err != nil {
		return nil, err
	}

	return result, nil
}

func UnmarshalItems[T any](items []map[string]types.AttributeValue) ([]T, error) {
	result := make([]T, len(items))

	for i, item := range items {
		if err := attributevalue.UnmarshalMapWithOptions(item, &result[i],
			func(options *attributevalue.DecoderOptions) {
				options.TagKey = "json"
			}); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func MarshalItem(item interface{}) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMapWithOptions(item,
		func(options *attributevalue.EncoderOptions) {
			options.TagKey = "json"
		},
	)

	if err != nil {
		return nil, err
	}

	return av, nil
}
