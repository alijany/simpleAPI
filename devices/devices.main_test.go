package devices

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	handler, err := Handler(mockConfig)

	if err != nil {
		t.Error("error in putDevice", err)
	}
	// check put request
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
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(body),
	}
	res, err := handler(request)
	if err != nil && res.StatusCode != http.StatusCreated {
		t.Error("error in put : ", err, res)
	}
	// check get request
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "api/devices/id1",
		QueryStringParameters: map[string]string{
			"id": "/devices/id1",
		},
	}
	res, err = handler(request)
	if err != nil || res.StatusCode != http.StatusOK {
		t.Error("error in get", err, res)
	}
	// check not allowed method
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPut,
		Path:       "api/devices/id1",
		QueryStringParameters: map[string]string{
			"id": "/devices/id1",
		},
	}
	res, err = handler(request)
	if err != nil || res.StatusCode != http.StatusMethodNotAllowed {
		t.Error("error in get", err, res)
	}
}
