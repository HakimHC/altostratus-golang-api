package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func initDynamoDb() {
	// TODO: get from environment vars
	region := "eu-west-1"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalf("Failed to start AWS session: %v", err)
	}

	DynamoDB = dynamodb.New(sess)
}
