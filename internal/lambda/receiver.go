package lambda

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleReceiveRequest(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	for _, message := range sqsEvent.Records {
		var myStruct MyStruct
		err := json.Unmarshal([]byte(message.Body), &myStruct)
		if err != nil {
			log.Println("Error unmarshalling message,", err)
			return "", err
		}

		log.Printf("Received message: %+v\n", myStruct)
	}

	return "Successfully processed SQS event", nil
}

func RunLambdaReceiver() error {
	lambda.Start(HandleReceiveRequest)
	return nil
}
