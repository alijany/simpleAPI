package devices

import (
	"main/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func putDevice(config *common.Config, device *Device) error {
	// check if DeviceModel exist
	_, err := getModel(config, device.DeviceModel)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(deviceTableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Id":          {S: aws.String(device.Id)},
			"DeviceModel": {S: aws.String(device.DeviceModel)},
			"Name":        {S: aws.String(device.Name)},
			"Note":        {S: aws.String(device.Note)},
			"Serial":      {S: aws.String(device.Serial)},
		},
	}

	_, err = config.DynamoDB.PutItem(input)
	return err
}
