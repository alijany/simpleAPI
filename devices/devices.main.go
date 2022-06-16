package devices

import (
	"main/common"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type HandlerCallback func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func Handler(config *common.Config) (HandlerCallback, error) {
	// Create tables (concurrent)
	channel := make(chan error, 2)
	go createDeviceTable(config, channel)
	go createModelTable(config, channel)

	// skip errors only if tables exist!
	for i := 0; i < 2; i++ {
		err := <-channel
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() != dynamodb.ErrCodeTableAlreadyExistsException {
			return nil, err
		}
	}

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
