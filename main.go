package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)

var snsClient *SNS

func initSNSClient() {
	var err error
	snsClient, err = NewSNSClient()
	if err != nil {
		log.Errorf("error while creating SNS client: %v", err)
	}
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	initSNSClient()
}

func handler(_ context.Context) error {
	m := []snsMessage{
		{
			AccountID: "1",
			ObjectID:  "1",
			Event:     "insert",
			Object:    "message",
			Origin:    "lambda",
		},
	}

	snsClient.Publish(m, os.Getenv("SNS_MESSAGE_TOPIC_ARN"))
	return nil
}

func main() {
	lambda.Start(handler)
}
