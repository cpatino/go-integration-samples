package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	awsRegion := "eu-west-1"
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	awsEndpoint := "http://localhost:4566"
	sqsClient := sqs.NewFromConfig(awsCfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(awsEndpoint)
	})

	queueURL := "http://localhost:4566/000000000000/my-queue"
	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 1,
	}

	resp, err := sqsClient.ReceiveMessage(context.TODO(), receiveMessageInput)
	if err != nil {
		log.Fatalf("Cannot receive messages: %s", err)
	}
	log.Printf("Received messages: %v", resp.Messages)
}
