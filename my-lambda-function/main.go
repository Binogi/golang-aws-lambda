package main

import (
	"encoding/json"
	"example.com/my-lambda-function/api_gateway" // "example.com/my-lambda-function" is the module name in go.mod
	"fmt"
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

	// Print request object
	reqData, _ := json.Marshal(req)
	fmt.Println("req object:", string(reqData))

	response := ExampleResponse{Message: fmt.Sprintf("Hello! Dear %s!", body.Name)}
	return api_gateway.ApiResponse(200, response)
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
