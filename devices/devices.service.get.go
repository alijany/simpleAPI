package devices

import (
	"encoding/json"
	"fmt"
	"main/common"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
)

var idRegex = regexp.MustCompile(`^/devices/id\d+$`)

func get(config *common.Config, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	if !idRegex.MatchString(id) {
		return common.ClientErrorMsg(
			http.StatusBadRequest,
			fmt.Sprintf(`{"message":"invalid id: %v"}`, id),
		)
	}

	device, err := getItemById[Device](config, deviceTableName, id)
	if err != nil {
		return common.ServerError(err)
	}
	if device == nil {
		return common.ClientError(http.StatusNotFound)
	}

	js, err := json.Marshal(device)
	if err != nil {
		return common.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}
