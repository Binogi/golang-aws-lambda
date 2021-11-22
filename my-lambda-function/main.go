package main

import (
	"example.com/my-lambda-function/api_gateway" // "example.com/my-lambda-function" is the module name in go.mod
	"fmt"
  "encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BodyParameters struct {
  Name string `json:"name"`
}

type ExampleResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
  body := BodyParameters{}
  json.Unmarshal([]byte(req.Body), &body)
	return api_gateway.ApiResponse(200, ExampleResponse{Message: fmt.Sprintf("Hello! Dear %s!", body.Name)})
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
