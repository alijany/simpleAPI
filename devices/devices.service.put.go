package devices

import (
	"encoding/json"
	"main/common"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func put(config *common.Config, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" &&
		req.Headers["Content-Type"] != "application/json" {
		return common.ClientError(http.StatusNotAcceptable)
	}

	device := new(Device)
	err := json.Unmarshal([]byte(req.Body), device)
	if err != nil {
		return common.ClientError(http.StatusUnprocessableEntity)
	}

	if device.Name == "" || device.Note == "" ||
		device.Serial == "" || device.DeviceModel == "" ||
		!idRegex.MatchString(device.DeviceModel) ||
		!idRegex.MatchString(device.Id) {
		return common.ClientError(http.StatusBadRequest)
	}

	err = putItem(config, device)
	if err != nil {
		return common.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}
