package devices

import (
	"main/common"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamodbMockClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *DynamodbMockClient) GetItem(input *dynamodb.GetItemInput) (
	*dynamodb.GetItemOutput,
	error,
) {
	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id":          {S: aws.String("/devices/id1")},
			"DeviceModel": {S: aws.String("/devicemodels/id1")},
			"Name":        {S: aws.String("Sensor")},
			"Note":        {S: aws.String("Testing a sensor.")},
			"Serial":      {S: aws.String("A020000102")},
		},
	}, nil
}

func TestGetDevice(t *testing.T) {
	mockSvc := &common.Config{DynamoDB: &DynamodbMockClient{}}

	device, err := getDevice(mockSvc, "/devices/id1")

	if err != nil || device == nil {
		t.Error("error in getDevice", err, device)
	}
}
