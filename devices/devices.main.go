package devices

import (
	"main/common"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type HandlerCallback func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func Handler(config *common.Config) (HandlerCallback, error) {
	// Create tables (concurrent)
	channel := make(chan error, 2)
	go createDeviceTable(config, channel)
	go createModelTable(config, channel)
	// TODO skip errors only if tables exist!
	_, _ = <-channel, <-channel //

	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.HTTPMethod {
		case "GET":
			return get(config, req)
		case "POST":
			return put(config, req)
		default:
			return common.ClientError(http.StatusMethodNotAllowed)
		}
	}, nil
}
