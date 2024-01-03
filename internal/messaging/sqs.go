package messaging

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type SQSMessage struct {
}

func NewSQSMessage() *SQSMessage {
	return &SQSMessage{}
}

func (m *SQSMessage) Publish(data interface{}) error {
	sqsURL := "https://sqs.ap-northeast-2.amazonaws.com/358059338173/falcon-worker-sqs"
	region := "ap-northeast-2"
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		log.Println("Error creating session,", err)
		return err
	}

	svc := sqs.New(sess)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling struct,", err)
		return err
	}

	jsonStr := string(jsonData)

	params := &sqs.SendMessageInput{
		MessageBody: aws.String(jsonStr),
		QueueUrl:    aws.String(sqsURL),
		//DelaySeconds: aws.Int64(1),
	}

	_, err = svc.SendMessage(params)
	if err != nil {
		log.Println("Error", err)
		return err
	}

	log.Println("Successfully sent message to SQS queue")
	return nil
}
