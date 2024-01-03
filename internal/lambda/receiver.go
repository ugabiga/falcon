package lambda

import (
	"context"
	"encoding/json"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/messaging"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleReceiveRequest(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	a := app.InitializeApplication()
	for _, message := range sqsEvent.Records {
		var reqData messaging.DCAMessage
		err := json.Unmarshal([]byte(message.Body), &reqData)
		if err != nil {
			log.Println("Error unmarshalling message,", err)
			return "", err
		}

		log.Printf("Received message: %+v\n", reqData)
		if err := a.RunLambdaSQS(reqData); err != nil {
			return "", err
		}
	}

	return "Successfully processed SQS event", nil
}

func RunLambdaReceiver() error {
	lambda.Start(HandleReceiveRequest)
	return nil
}
