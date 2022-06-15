package main

import (
	"main/common"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(`{ "hello": "mamad" }`),
		}, nil
	case "POST":
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(`{ "hello": "mamad" }`),
		}, nil
	default:
		return common.ClientError(http.StatusMethodNotAllowed)
	}
}

func main() {
	lambda.Start(handler)
}
