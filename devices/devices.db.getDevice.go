package devices

import (
	"main/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getDevice(config *common.Config, id string) (*Device, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String(deviceTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	}

	// Retrieve the item from DynamoDB.
	result, err := config.DynamoDB.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	device := new(Device)
	err = dynamodbattribute.UnmarshalMap(result.Item, device)
	if err != nil {
		return nil, err
	}

	return device, nil
}
