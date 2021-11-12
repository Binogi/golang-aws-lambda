package main

import (
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
)

type ExampleEvent struct {
  Name string `json:"name"`
}

type ExampleResponse struct {
  Message string `json:"message"`
}

func HandleLambdaEvent(event ExampleEvent) (ExampleResponse, error) {
  return ExampleResponse{Message: fmt.Sprintf("Hello! Dear %s!", event.Name)}, nil
}

func main() {
  lambda.Start(HandleLambdaEvent)
}
