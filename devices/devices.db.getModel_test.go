package devices

import (
	"main/common"
	"testing"
)

func TestGetModel(t *testing.T) {
	mockSvc := &common.Config{DynamoDB: &DynamodbMockClient{}}

	device, err := getDevice(mockSvc, "/devices/id1")

	if err != nil || device == nil {
		t.Error("error in getDevice", err, device)
	}
}
