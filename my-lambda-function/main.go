package main

import (
	"example.com/my-lambda-function/api_gateway" // "example.com/my-lambda-function" is the module name in go.mod
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ExampleResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return api_gateway.ApiResponse(200, ExampleResponse{Message: fmt.Sprintf("Hello! Dear %s!", req.QueryStringParameters["name"])})
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
