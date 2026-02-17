package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const QueueURL = "http://localhost:4566/000000000000/my-queue"

func main() {
	sqsClient := GetSqsClient()
	receiveMessages(sqsClient)
}

func receiveMessages(sqsClient *sqs.Client) {
	for {
		results, err := sqsClient.ReceiveMessage(
			context.TODO(),
			&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(QueueURL),
				MaxNumberOfMessages: 10,
			})

		if err != nil {
			log.Fatal(err)
		}

		for _, message := range results.Messages {
			processMessage(message, sqsClient)
		}
	}
}

func processMessage(message types.Message, sqsClient *sqs.Client) {

	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(QueueURL),
		ReceiptHandle: message.ReceiptHandle,
	}

	log.Printf("Received message: %s", *message.Body)

	_, err := sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}
