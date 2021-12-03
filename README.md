# Go functions on AWS Lambda

[![Develop in Go, and help kids learn better](https://media.cdn.teamtailor.com/images/s3/teamtailor-production/gallery_picture/image_uploads/90c2c459-c30e-4f30-aea3-b7f5e47f52be/original.png)](https://jobs.binogi.com/jobs?department=Tech)

Uses the following technologies:

- **Go** programming language
- **AWS Lambda** to host the function
- **AWS API Gateway** to expose the Lambda function as a REST API*
- **AWS CLI tool** (`aws`) to configure AWS

The current setup in this project is: 1 REST endpoint → 1 Lambda function → inside 1 ZIP file.

### Notes on AWS API Gateway

*NOTE: AWS API Gateway adds a bit of overhead on how you handle parameters (`APIGatewayProxyRequest`) and response (`APIGatewayProxyResponse`), which affects how you test it on localhost.



## Inspiration and references

Based on [AWS Lambda in GoLang — The Ultimate Guide](https://www.softkraft.co/aws-lambda-in-golang/) and [this example Gist](https://gist.github.com/josephspurrier/05b9126279703a81122cba198df50d6f). See also:

- AWS:
	- [AWS Go guide](https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html)
	- [Example of AWS Lambda function written in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
	- [AWS API Gateway tutorial: turn Lambda function into REST API](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html)


## Write your Lambda function in Go

Edit `my-lambda-function/main.go`


## Building both Go binary and ZIP file

    make


### (Optional: Manual build steps for Go binary and ZIP file)

#### Build Go binary

Create the working directory

    mkdir my-lambda-function && cd my-lambda-function

_(Create the files: `main.go`)_

Initialize the module in Go

    go mod init example.com/my-lambda-function

Build the go app for Linux

    GOOS=linux go build

#### Build ZIP file

To build your ZIP file:

    cd my-lambda-function && zip -r ../my-lambda-function.zip * && cd ..


## Running the Lambda function locally

**TBD!**

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"body":{"name":"Tom"}}'

✅ You should get a response like `{"statusCode":200,"headers":{"Content-Type":"application/json"},"multiValueHeaders":null,"body":"{\"message\":\"Hello! Dear Tom!\"}"}`

Note the more complex JSON parameters/response used vs in the final results below, see [Notes on AWS API Gateway](#notes-on-aws-api-gateway) above.


## Deploy to AWS Lambda

Parameters used below:

- `[my-lambda-function]`: substitute with your app/repository/image name
- `[AWS Account Number]`: e.g. `123456789012`
- `[Region]`: e.g. `eu-west-1`, `us-east-1`

You need an IAM Role (e.g. `lambda-ex`) for _executing_ the Lambda function.

Build your local ZIP file:

    make

### Create a new Lambda function

    aws lambda create-function --function-name [my-lambda-function] \
      --handler [my-lambda-function] \
      --zip-file fileb://tmp/[my-lambda-function].zip \
      --runtime go1.x \
      --region [Region] \
      --role arn:aws:iam::[AWS Account Number]:role/lambda-execute

✅ You can now test your Lambda function [in the AWS Console, “Test” tab](https://console.aws.amazon.com/lambda/home) or with:

    aws lambda invoke --function-name [my-lambda-function] \
      --cli-binary-format raw-in-base64-out \
      --payload '{"body":"{\"name\":\"Tom\"}"}' \
      --invocation-type "RequestResponse" tmp/lambda_response.txt

### Turn the Lambda function into a REST API using AWS API Gateway

- Go to https://console.aws.amazon.com/apigateway
- REST API → Build
- Select 🔘 “New API”
- Actions → Create Resource, enter resource name e.g. `my-lambda-function`
- Select your created resource, then Actions → Create Method
- Select method (e.g. `POST`, or `ANY`)
- Select Integration type: 🔘 “Lambda Function”, then check the ☑️ “Use Lambda Proxy integration” and enter the name of your function under “Lambda Function”
- Actions → Deploy API
- Enter a “Stage” name e.g. “test”

You will get an endpoint back, e.g. `https://12345xw4tf.execute-api.eu-west-1.amazonaws.com/test` (note: excludes the function name, see the full `curl` example below)

✅ You can now test the API with:

    curl -XPOST "https://12345xw4tf.execute-api.eu-west-1.amazonaws.com/test/my-lambda-function" -d '{"name":"Tom"}'

### Update an existing Lambda function

Update the Lambda function:

    aws lambda update-function-code --region [Region] --function-name [my-lambda-function] \
      --zip-file fileb://tmp/[my-lambda-function].zip

**Note:** There’s no need to update API Gateway after updating the Lambda function, but **it can take a minute or so** before the updated Lambda function starts working.


## Troubleshooting

- Note that your Lambda function and API Gateway will have **separate** *Log Groups* in [CloudWatch](https://console.aws.amazon.com/cloudwatch/home).
- If you get API Gateway "Internal server error", "malformed Lambda proxy response" errors or 502 status codes, it means [your Lambda function’s response is not formatted correctly for API Gateway](https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/).
