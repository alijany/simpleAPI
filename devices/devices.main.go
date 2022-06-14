package devices

import (
	"main/common"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(config *common.Config) {
	// Create tables (concurrent)
	// TODO: skip errors only if tables exist!
	channel := make(chan error, 2)
	go createDeviceTable(config, channel)
	go createModelTable(config, channel)
	_, _ = <-channel, <-channel

	router := func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.HTTPMethod {
		case "GET":
			return get(config, req)
		case "POST":
			return put(config, req)
		default:
			return common.ClientError(http.StatusMethodNotAllowed)
		}
	}

	lambda.Start(router)
}
