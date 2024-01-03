package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func HandleRequest(ctx context.Context) (string, error) {
	log.Println("HandleRequest")
	log.Printf("ctx: %+v", ctx)
	//a := app.InitializeApplication()
	//if err := a.RunLambdaServer(); err != nil {
	//	return "", err
	//}

	return "ok", nil
}

func RunLambdaWorker() error {
	lambda.Start(HandleRequest)

	return nil
}
