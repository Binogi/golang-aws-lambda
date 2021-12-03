package main

import (
	"encoding/json"
	"example.com/my-lambda-function/api_gateway" // "example.com/my-lambda-function" is the module name in go.mod
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type BodyParameters struct {
	Name string `json:"name"`
}

type ExampleResponse struct {
	Message string `json:"message"`
}

func HandleLocalEvent(res http.ResponseWriter, req *http.Request) {
	// Print request object
	reqData, _ := json.Marshal(req)
	fmt.Println("req object:", req.Body, string(reqData))
	fmt.Fprintf(res, "Hello world")
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
	isLocal := flag.Bool("local", false, "Set this flag if you are running locally (not on Lambda)")
	flag.Parse()

	if *isLocal {
		fmt.Print("--local flag has been set")
		http.HandleFunc("/", HandleLocalEvent)
		http.ListenAndServe(":3000", nil)
	} else {
		lambda.Start(HandleLambdaEvent)
	}
}
