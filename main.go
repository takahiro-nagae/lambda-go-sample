package main

import (
	"lambda-go-sample/src"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(src.HandleRequest)
}
