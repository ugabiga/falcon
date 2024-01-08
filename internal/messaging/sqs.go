package messaging

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/service"
	"log"
)

type SQSMessageHandler struct {
	dcaSrv *service.DcaService
}

func NewSQSMessageHandler(
	dcaSrv *service.DcaService,
) *SQSMessageHandler {
	return &SQSMessageHandler{
		dcaSrv: dcaSrv,
	}
}

func (h SQSMessageHandler) Publish() error {
	lambda.Start(h.HandleReceiveRequest)
	return nil
}

func (h SQSMessageHandler) HandlePublish(ctx context.Context) (string, error) {
	log.Printf("Start watching messages from DCA")

	messages, err := h.dcaMessages()
	if err != nil {
		log.Printf("Error occurred during watching. Err: %v", err)
		return "", err
	}
	log.Printf("DCA messages count: %d", len(messages))
	log.Printf("DCA messages: %+v", debug.ToJSONStr(messages))

	log.Printf("Start publishing messages to SQS")
	for _, msg := range messages {
		if err := h.publish(msg); err != nil {
			log.Printf("Error occurred during publishing. Err: %v", err)
			continue
		}
	}

	return "ok", nil
}

func (h SQSMessageHandler) publish(data interface{}) error {
	log.Printf("Publishing message to SQS: %+v", data)
	log.Printf("Publishing message to SQS: %+v", debug.ToJSONStr(data))

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

	log.Printf("jsonStr: %s", jsonStr)

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

func (h SQSMessageHandler) Subscribe() error {
	lambda.Start(h.HandleReceiveRequest)
	return nil
}

func (h SQSMessageHandler) HandleReceiveRequest(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	for _, message := range sqsEvent.Records {
		var reqData DCAMessage
		err := json.Unmarshal([]byte(message.Body), &reqData)
		if err != nil {
			log.Println("Error unmarshalling message,", err)
			continue
		}

		log.Printf("Received message: %+v\n", reqData)
		if err := h.subscribe(reqData); err != nil {
			log.Println("Error subscribing message,", err)
			continue
		}
	}

	return "Successfully processed SQS event", nil
}

func (h SQSMessageHandler) subscribe(reqMsg DCAMessage) error {
	if err := h.dcaSrv.Order(reqMsg.TaskOrderInfo); err != nil {
		log.Printf("Error occurred during order. Err: %v", err)
		return err
	}
	return nil
}

func (h SQSMessageHandler) dcaMessages() ([]DCAMessage, error) {
	messages, err := h.dcaSrv.GetTarget()
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d messages", len(messages))
	log.Printf("Found messages: %+v", debug.ToJSONStr(messages))

	var dcaMessages []DCAMessage
	for _, msg := range messages {
		log.Printf("Found message: %+v", debug.ToJSONStr(msg))
		dcaMessages = append(dcaMessages, DCAMessage{
			TaskOrderInfo: msg,
		})
	}

	return dcaMessages, nil
}
