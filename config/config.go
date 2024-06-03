package config

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

var DynamoDB *dynamodb.DynamoDB
var JwtSecret string

func init() {
	initDynamoDb()
	initConfigVars()
}

func initConfigVars() {
	JwtSecret = os.Getenv("JWT_SECRET")
}
