package lambda

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func Run() {
	sqsURL := ""
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		log.Println("Error creating session,", err)
		return
	}

	svc := sqs.New(sess)

	myStruct := MyStruct{
		Field1: "Test",
		Field2: 123,
	}

	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		log.Println("Error marshalling struct,", err)
		return
	}

	jsonStr := string(jsonData)

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(jsonStr),
		QueueUrl:     aws.String(sqsURL),
		DelaySeconds: aws.Int64(1),
	}

	_, err = svc.SendMessage(params)
	if err != nil {
		log.Println("Error", err)
		return
	}

	log.Println("Successfully sent message to SQS queue")
}

func HandleCronRequest(ctx context.Context) (string, error) {
	log.Printf("HandleRequest: %v", ctx)

	Run()

	return "ok", nil
}

func RunLambdaCron() error {
	lambda.Start(HandleCronRequest)

	return nil
}
