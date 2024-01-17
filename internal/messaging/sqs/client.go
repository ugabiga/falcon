package sqs

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type Client struct {
	sqsURL    string
	sqsRegion string
}

func NewClient(sqsURL, sqsRegion string) *Client {
	return &Client{
		sqsURL:    sqsURL,
		sqsRegion: sqsRegion,
	}
}

func (c Client) Publish(data interface{}) (*sqs.SendMessageOutput, error) {
	sqsURL := c.sqsURL
	region := c.sqsRegion

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Println("Error creating session,", err)
		return nil, err
	}

	svc := sqs.New(sess)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling struct,", err)
		return nil, err
	}

	jsonStr := string(jsonData)

	params := &sqs.SendMessageInput{
		MessageBody: aws.String(jsonStr),
		QueueUrl:    aws.String(sqsURL),
	}

	output, err := svc.SendMessage(params)
	if err != nil {
		log.Println("Error sending message to SQS queue,", err)
		return nil, err
	}
	log.Printf("Message sent to SQS queue: %s", output)

	return output, nil
}

func (c Client) Pull() (*sqs.ReceiveMessageOutput, error) {
	sqsURL := c.sqsURL
	region := c.sqsRegion

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Println("Error creating session,", err)
		return nil, err
	}

	svc := sqs.New(sess)

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(sqsURL),
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(60),
	}
	sqsMessage, err := svc.ReceiveMessage(params)
	if err != nil {
		log.Printf("Error receiving message from SQS queue: %s", err)
		return nil, err
	}

	return sqsMessage, nil
}

func (c Client) Delete(message *sqs.Message) error {
	sqsURL := c.sqsURL
	region := c.sqsRegion

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Println("Error creating session,", err)
		return err
	}

	svc := sqs.New(sess)

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(sqsURL),
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err = svc.DeleteMessage(params)
	if err != nil {
		log.Printf("Error deleting message from SQS queue: %s", err)
		return err
	}

	return nil
}
