package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type ExampleEvent struct {
  Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name ExampleEvent) (string, error) {
  return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
  lambda.Start(HandleRequest)
}
