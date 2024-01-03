package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ugabiga/falcon/internal/app"
)

func HandleRequest(ctx context.Context) (string, error) {
	a := app.InitializeApplication()
	if err := a.RunLambdaServer(); err != nil {
		return "", err
	}

	return "ok", nil
}

func RunLambdaWorker() error {
	lambda.Start(HandleRequest)

	return nil
}
