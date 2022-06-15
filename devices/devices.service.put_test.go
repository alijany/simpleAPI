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
	// check normal request
	request := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "api/devices",
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(body),
	}
	res, err := put(mockConfig, request)
	if err != nil && res.StatusCode != http.StatusCreated {
		t.Error("error in put : ", err, res)
	}
	// check missed input
	body, err = json.Marshal(map[string]string{
		"id":          "/devices/id1",
		"deviceModel": "/devicemodels/id1",
		"note":        "Testing a sensor.",
		"serial":      "A020000102",
	})
	if err != nil {
		t.Error("error in put : ", err)
	}
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "api/devices",
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(body),
	}
	res, err = put(mockConfig, request)
	if err != nil && res.StatusCode != http.StatusBadRequest {
		t.Error("error in put : ", err, res)
	}

	// check invalid body input
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "api/devices",
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string("{]"),
	}
	res, err = put(mockConfig, request)
	if err != nil && res.StatusCode != http.StatusUnprocessableEntity {
		t.Error("error in put : ", err, res)
	}

	// check missed header
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "api/devices",
		Body:       string(""),
	}
	res, err = put(mockConfig, request)
	if err != nil && res.StatusCode != http.StatusNotAcceptable {
		t.Error("error in put : ", err, res)
	}
}
