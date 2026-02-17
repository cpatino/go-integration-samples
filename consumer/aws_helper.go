package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const awsRegion = "eu-west-1"
const awsEndpoint = "http://localhost:4566"

func loadAwsConfig() aws.Config {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}
	return awsCfg
}

func GetSqsClient() *sqs.Client {
	awsCfg := loadAwsConfig()
	sqsClient := sqs.NewFromConfig(awsCfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(awsEndpoint)
	})
	return sqsClient
}
