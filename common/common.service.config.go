package common

import (
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitializeConfig() (*Config, error) {
	// sess, err := session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// })

	// sess, err := session.NewSession(&aws.Config{
	// 	Credentials: credentials.AnonymousCredentials,
	// 	Region:      aws.String("us-west-2")}, // TODO: os AWS_REGION
	// )

	// if err != nil {
	// 	return nil, err
	// }

	sess := func() *session.Session {
		// server is the mock server that simply writes a 200 status back to the client
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		return session.Must(session.NewSession(&aws.Config{
			DisableSSL:  aws.Bool(true),
			Credentials: credentials.AnonymousCredentials,
			Endpoint:    aws.String(server.URL),
			Region:      aws.String("us-west-2"),
		}))
	}()

	return &Config{
		DynamoDB: dynamodb.New(sess),
	}, nil
}
