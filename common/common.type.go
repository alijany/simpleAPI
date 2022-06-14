package common

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type Config struct {
	DynamoDB dynamodbiface.DynamoDBAPI
}
