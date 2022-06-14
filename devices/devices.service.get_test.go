package devices

import (
	"main/common"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGet(t *testing.T) {
	config := &common.Config{DynamoDB: &DynamodbMockClient{}}
	request := events.APIGatewayProxyRequest{Body: "Paul"}
	get(config, request)
}
