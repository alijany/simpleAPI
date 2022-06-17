package devices

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGet(t *testing.T) {
	// check normal request
	request := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "api/devices/id1",
		PathParameters: map[string]string{
			"id": "/devices/id1",
		},
	}
	res, err := get(mockConfig, request)
	if err != nil || res.StatusCode != http.StatusOK {
		t.Error("error in get", err, res)
	}
	// check invalid id format
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "api/devices/id1",
		PathParameters: map[string]string{
			"id": "id1",
		},
	}
	res, err = get(mockConfig, request)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Error("error in get", err, res)
	}
	// check device not found
	request = events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "api/devices//devices/id2",
		PathParameters: map[string]string{
			"id": "/devices/id2",
		},
	}
	res, err = get(mockConfig, request)
	if err != nil || res.StatusCode != http.StatusNotFound {
		t.Error("error in get", err, res)
	}
}
