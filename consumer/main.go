package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const queueURL = "http://localhost:4566/000000000000/my-queue"

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

	for {
		results, err := sqsClient.ReceiveMessage(
			context.TODO(),
			&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(queueURL),
				MaxNumberOfMessages: 10,
			})

		if err != nil {
			log.Fatal(err)
		}

		for _, message := range results.Messages {
			processJob(message, sqsClient)
		}
	}
}

func processJob(message types.Message, sqsClient *sqs.Client) {

	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: message.ReceiptHandle,
	}

	log.Printf("Received message: %s", *message.Body)

	_, err := sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}
