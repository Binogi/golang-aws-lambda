# Go functions on AWS Lambda in Docker

[![Develop in Go, and help kids learn better](https://media.cdn.teamtailor.com/images/s3/teamtailor-production/gallery_picture/image_uploads/90c2c459-c30e-4f30-aea3-b7f5e47f52be/original.png)](https://jobs.binogi.com/jobs?department=Tech)

Uses the following technologies:

- **Go** programming language
- **Docker** as container
- **AWS Lambda** to host the function
- **AWS ECR (Elastic Container Registry)** to host the Docker image
- **AWS API Gateway** to expose the Lambda function as a REST API
- **AWS CLI tool** (`aws`) to configure AWS

The current setup in this project is: 1 REST endpoint → 1 Lambda function → inside 1 Docker image.

Based on [Docker Hub: aws-lambda-go](https://hub.docker.com/r/amazon/aws-lambda-go) and [this example Gist](https://gist.github.com/josephspurrier/05b9126279703a81122cba198df50d6f). See also:

- AWS:
	- [AWS Deploy Go Lambda functions with Docker](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html)
	- [AWS Go guide](https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html)
	- [AWS Docker image on ECR (Elastic Container Registry)](https://gallery.ecr.aws/lambda/go)
	- [AWS Lambda on Docker](https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html)
	- [Example of AWS Lambda function written in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
	- [AWS API Gateway tutorial: turn Lambda function into REST API](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html)
- Docker:
	- [Building Go Docker images](https://docs.docker.com/language/golang/build-images/)
- Other:
	- [AWS Lambda in GoLang — The Ultimate Guide](https://www.softkraft.co/aws-lambda-in-golang/)


## Write your Lambda function in Go

Edit `my-lambda-function/main.go`


## Building both Go binary and Docker image

    make


### (Optional: Manual build steps for Go binary and Docker image)

#### Build Go binary

Create the working directory

    mkdir my-lambda-function

Change to the directory

    cd my-lambda-function

_(Create the files: `main.go` and `Dockerfile`)_

Initialize the module in Go

    go mod init example.com/my-lambda-function

Build the go app for Linux

    GOOS=linux go build

#### Build Docker image

To build your Docker image:

    docker build -t my-lambda-function .


## Running the Lambda function locally

To run your image locally:

    docker run -p 9000:8080 my-lambda-function

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"Tom"}'

✅ You should get a response like `{"message":"Hello! Dear Tom!"}`

## Deploy to AWS Lambda

Parameters used below:

- `[my-lambda-function]`: substitute with your app/repository/image name
- `[AWS Account Number]`: e.g. `123456789012`
- `[Region]`: e.g. `eu-west-1`, `us-east-1`

Create an IAM Policy (e.g. `ECRDockerImageCreation`) with the following permissions:

    ecr:*

You also need an IAM Role (e.g. `lambda-ex`) for _executing_ the Lambda function.

Build your local Docker image:

    make

### Create a new Lambda function

Authenticate the Docker CLI to your Amazon ECR registry:

    aws ecr get-login-password --region [Region] | docker login --username AWS --password-stdin [AWS Account Number].dkr.ecr.[Region].amazonaws.com

✅ You should get a response like `Login Succeeded`

Tag your new Docker image to match your repository name:

    docker tag [my-lambda-function]:latest [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

(You won’t get any response)

Create an ECR repository:

    aws ecr create-repository --repository-name [my-lambda-function] --region [Region]

Deploy the Docker image to Amazon ECR:

    docker push [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

You should now see your image repository on https://console.aws.amazon.com/ecr/repositories?region=[Region]

    aws lambda create-function --region [Region] --function-name [my-lambda-function] \
      --package-type Image \
      --code ImageUri=[AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest \
      --role arn:aws:iam::[AWS Account Number]:role/lambda-ex

✅ You can now test your Lambda function [in the AWS Console, “Test” tab](https://console.aws.amazon.com/lambda/home).

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

Tag your new Docker image to match your repository name:

    docker tag [my-lambda-function]:latest [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

Deploy the Docker image to Amazon ECR:

    docker push [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

Update the Lambda function:

    aws lambda update-function-code --region [Region] --function-name [my-lambda-function] \
      --image-uri=[AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

**Note:** There’s no need to update API Gateway after updating the Lambda function, but **it can take a minute or so** before the updated Lambda function starts working.

## Troubleshooting

- Note that your Lambda function and API Gateway will have **separate** *Log Groups* in [CloudWatch](https://console.aws.amazon.com/cloudwatch/home).
- If you get API Gateway "Internal server error", "malformed Lambda proxy response" errors or 502 status codes, it means [your Lambda function’s response is not formatted correctly for API Gateway](https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/).
