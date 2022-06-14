package devices

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestPut(t *testing.T) {
	body, err := json.Marshal(map[string]string{
		"id":          "/devices/id1",
		"deviceModel": "/devicemodels/id1",
		"name":        "Sensor",
		"note":        "Testing a sensor.",
		"serial":      "A020000102",
	})

	if err != nil {
		t.Error("error in put : ", err)
	}

	request := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "api/devices",
		Body:       string(body),
	}
	res, err := put(mockConfig, request)

	if err != nil {
		t.Error("error in put : ", err, res)
	}
}
