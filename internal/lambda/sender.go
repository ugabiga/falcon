package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ugabiga/falcon/internal/app"
)

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func HandleCronRequest(ctx context.Context) (string, error) {
	a := app.InitializeApplication()
	if err := a.RunLambdaCron(); err != nil {
		return "", err
	}
	return "ok", nil
}

func RunLambdaCron() error {
	lambda.Start(HandleCronRequest)

	return nil
}
