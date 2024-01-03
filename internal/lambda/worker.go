package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ugabiga/falcon/internal/app"
	"log"
)

func HandleRequest(ctx context.Context) (string, error) {
	log.Printf("HandleRequest: %v", ctx)

	a := app.InitializeApplication()
	if err := a.Worker(); err != nil {
		return "", err
	}

	return "ok", nil
}

func RunLambdaWorker() error {
	lambda.Start(HandleRequest)

	return nil
}
