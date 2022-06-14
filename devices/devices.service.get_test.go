package devices

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGet(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "api/devices/id1",
		QueryStringParameters: map[string]string{
			"id": "/devices/id1",
		},
	}
	res, err := get(mockConfig, request)

	if err != nil {
		t.Error("error in get", err, res)
	}
}
