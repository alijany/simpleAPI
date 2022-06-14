package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitializeConfig() (*Config, error) {
	// sess, err := session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// })

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("us-west-2")}, // TODO: os AWS_REGION
	)

	if err != nil {
		return nil, err
	}

	return &Config{
		DynamoDB: dynamodb.New(sess),
	}, nil
}
