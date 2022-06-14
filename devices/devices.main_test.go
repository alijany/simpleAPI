package devices

import (
	"main/common"
	"testing"
)

func TestHandler(t *testing.T) {
	config := &common.Config{DynamoDB: &DynamodbMockClient{}}

	Handler(config)
}
