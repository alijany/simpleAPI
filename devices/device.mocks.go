package devices

import (
	"main/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamodbMockClient struct {
	dynamodbiface.DynamoDBAPI
}

var mockConfig = &common.Config{DynamoDB: &DynamodbMockClient{}}

// overwrite DynamoDBAPI
func (m *DynamodbMockClient) GetItem(input *dynamodb.GetItemInput) (
	*dynamodb.GetItemOutput,
	error,
) {
	switch *input.TableName {
	case deviceTableName:
		return &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"Id":          {S: aws.String("/devices/id1")},
				"DeviceModel": {S: aws.String("/devicemodels/id1")},
				"Name":        {S: aws.String("Sensor")},
				"Note":        {S: aws.String("Testing a sensor.")},
				"Serial":      {S: aws.String("A020000102")},
			},
		}, nil
	case modelTableName:
		return &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"Id":   {S: aws.String("/devicemodels/id1")},
				"Name": {S: aws.String("Test model")},
				"Note": {S: aws.String("Testing a model.")},
			},
		}, nil
	default:
		return &dynamodb.GetItemOutput{}, &dynamodb.TableNotFoundException{}
	}
}

func (m *DynamodbMockClient) PutItem(input *dynamodb.PutItemInput) (
	*dynamodb.PutItemOutput,
	error,
) {
	return &dynamodb.PutItemOutput{
		Attributes: map[string]*dynamodb.AttributeValue{
			"Id":          {S: aws.String("/devices/id1")},
			"DeviceModel": {S: aws.String("/devicemodels/id1")},
			"Name":        {S: aws.String("Sensor")},
			"Note":        {S: aws.String("Testing a sensor.")},
			"Serial":      {S: aws.String("A020000102")},
		},
	}, nil
}

func (m *DynamodbMockClient) CreateTable(*dynamodb.CreateTableInput) (
	*dynamodb.CreateTableOutput,
	error,
) {
	return &dynamodb.CreateTableOutput{}, nil
}
