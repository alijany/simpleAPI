package main

import (
	"log"
	"main/common"
	"main/devices"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config, err := common.InitializeConfig()
	if err != nil {
		log.Fatal(err)
	}

	router, err := devices.Handler(config)
	if err != nil {
		log.Fatal(err)
	}
	lambda.Start(router)
}
