package common

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func ClientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
