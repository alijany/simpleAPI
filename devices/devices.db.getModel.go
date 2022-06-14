package devices

import (
	"main/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getModel(config *common.Config, id string) (*DeviceModel, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String(modelTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	}

	// Retrieve the item from DynamoDB.
	// If no matching item is found return nil.
	result, err := config.DynamoDB.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	model := new(DeviceModel)
	err = dynamodbattribute.UnmarshalMap(result.Item, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
