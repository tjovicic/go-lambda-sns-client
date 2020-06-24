package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type SNS struct {
	Client *sns.SNS
}

func NewSNSClient() (*SNS, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "error while creating aws session for sns client")
	}

	return &SNS{Client: sns.New(sess)}, nil
}

func (c *SNS) Publish(records []snsMessage, topicArn string) {
	for _, r := range records {
		b, err := json.Marshal(r)
		if err != nil {
			log.Errorf("%+v", errors.Wrap(err, "error while marshalling "))
			continue
		}

		_, err = c.Client.Publish(&sns.PublishInput{
			Message:  aws.String(string(b)),
			TopicArn: &topicArn,
		})
		if err != nil {
			log.Errorf("%+v", errors.Wrap(err, "error while publishing to SNS"))
		}
	}
}

type snsMessage struct {
	AccountID string `json:"account_id"`
	ObjectID  string `json:"object_id"`
	Event     string `json:"event"`
	Object    string `json:"object"`
	Origin    string `json:"origin"`
}
