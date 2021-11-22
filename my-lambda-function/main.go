package main

import (
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"
  "example.com/my-lambda-function/api_gateway" // "example.com/my-lambda-function" is the module name in go.mod
)

type ExampleEvent struct {
  Name string `json:"name"`
}

type ExampleResponse struct {
  Message string `json:"message"`
}

func HandleLambdaEvent(event ExampleEvent) (*events.APIGatewayProxyResponse, error) {
  return api_gateway.ApiResponse(200, ExampleResponse{Message: fmt.Sprintf("Hello! Dear %s!", event.Name)})
}

func main() {
  lambda.Start(HandleLambdaEvent)
}
