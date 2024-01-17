package messaging

import (
	"context"
	"encoding/json"
	"github.com/AlekSi/pointer"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
)

type SQSMessageHandler struct {
	sqsURL              string
	sqsRegion           string
	sqsSubscriptionType string
	dcaSrv              *service.DcaService
	gridSrv             *service.GridService
}

func NewSQSMessageHandler(
	cfg *config.Config,
	dcaSrv *service.DcaService,
	gridSrv *service.GridService,
) *SQSMessageHandler {
	return &SQSMessageHandler{
		sqsURL:              cfg.SQSQueueURL,
		sqsRegion:           cfg.AWSRegion,
		sqsSubscriptionType: cfg.SQSSubscriptionType,
		dcaSrv:              dcaSrv,
		gridSrv:             gridSrv,
	}
}

func (h SQSMessageHandler) Publish() error {
	lambda.Start(h.HandlePublish)
	return nil
}

func (h SQSMessageHandler) Subscribe() error {
	if h.sqsSubscriptionType == "pull" {

	}
	lambda.Start(h.HandleReceiveRequest)
	return nil
}

func (h SQSMessageHandler) dcaMessages() ([]TaskOrderInfoMessage, error) {
	messages, err := h.dcaSrv.GetTarget()
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d messages", len(messages))
	log.Printf("Found messages: %+v", debug.ToJSONInlineStr(messages))

	var dcaMessages []TaskOrderInfoMessage
	for _, msg := range messages {
		log.Printf("Found message: %+v", debug.ToJSONInlineStr(msg))
		dcaMessages = append(dcaMessages, TaskOrderInfoMessage{
			TaskOrderInfo: msg,
		})
	}

	return dcaMessages, nil
}

func (h SQSMessageHandler) gridMessages() ([]TaskOrderInfoMessage, error) {
	messages, err := h.gridSrv.GetTarget()
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d messages", len(messages))
	log.Printf("Found messages: %+v", debug.ToJSONInlineStr(messages))

	var gridMessages []TaskOrderInfoMessage
	for _, msg := range messages {
		log.Printf("Found message: %+v", debug.ToJSONInlineStr(msg))
		gridMessages = append(gridMessages, TaskOrderInfoMessage{
			TaskOrderInfo: msg,
		})
	}

	return gridMessages, nil

}

func (h SQSMessageHandler) HandlePublish(ctx context.Context) (string, error) {
	log.Printf("Start publishing messages to SQS")

	dcaMessages, err := h.dcaMessages()
	if err != nil {
		log.Printf("Error occurred during getting messages. Err: %v", err)
		return "", err
	}
	log.Printf("DCA messages count: %d", len(dcaMessages))

	for _, msg := range dcaMessages {
		if err := h.publishMessage(msg); err != nil {
			log.Printf("Error occurred during publishing. Err: %v", err)
			continue
		}
	}

	//gridMessages, err := h.gridMessages()
	//if err != nil {
	//	log.Printf("Error occurred during getting messages. Err: %v", err)
	//	return "", err
	//}
	//log.Printf("Grid messages count: %d", len(gridMessages))
	//
	//for _, msg := range gridMessages {
	//	if err := h.publishMessage(msg); err != nil {
	//		log.Printf("Error occurred during publishing. Err: %v", err)
	//		continue
	//	}
	//}

	return "ok", nil
}

func (h SQSMessageHandler) subscribe(reqMsg TaskOrderInfoMessage) error {
	if err := h.dcaSrv.Order(reqMsg.TaskOrderInfo); err != nil {
		log.Printf("Error occurred during order. Err: %v", err)
		return err
	}
	return nil
}

func (h SQSMessageHandler) HandleReceiveRequest(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	for _, message := range sqsEvent.Records {
		var reqData TaskOrderInfoMessage
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

func (h SQSMessageHandler) pullToSubscribeMessages() error {
	sqsURL := h.sqsURL
	region := h.sqsRegion

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Println("Error creating session,", err)
		return err
	}

	svc := sqs.New(sess)

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(sqsURL),
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(60),
	}
	sqsMessage, err := svc.ReceiveMessage(params)
	if err != nil {
		return err
	}

	for _, msg := range sqsMessage.Messages {
		var reqData TaskOrderInfoMessage
		if err := json.Unmarshal([]byte(pointer.GetString(msg.Body)), &reqData); err != nil {
			log.Println("Error unmarshalling message,", err)
			continue
		}

		log.Printf("Received message: %+v\n", debug.ToJSONStr(reqData))
	}

	return nil
}

func (h SQSMessageHandler) publishMessage(data interface{}) error {
	sqsURL := h.sqsURL
	region := h.sqsRegion

	log.Printf("Publishing message to SQS: %+v", debug.ToJSONInlineStr(data))

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
	}

	_, err = svc.SendMessage(params)
	if err != nil {
		log.Println("Error sending message to SQS queue,", err)
		return err
	}

	log.Println("Successfully sent message to SQS queue")
	return nil
}
