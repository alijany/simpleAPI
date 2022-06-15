package common

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitializeConfig() (*Config, error) {
	// get region
	region, ok := os.LookupEnv("AWS_REGION")
	if ok {
		region = "us-east-1"
	}
	// create session
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		return nil, err
	}
	// return config
	return &Config{
		DynamoDB: dynamodb.New(sess),
	}, nil
}
