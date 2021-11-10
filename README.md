# Go functions on AWS Lambda in Docker

Based on [aws-lambda-go on Docker Hub](https://hub.docker.com/r/amazon/aws-lambda-go) and [this example Gist](https://gist.github.com/josephspurrier/05b9126279703a81122cba198df50d6f).

See also:

- AWS Deploy Go Lambda functions with Docker: https://docs.aws.amazon.com/lambda/latest/dg/go-image.html
- AWS Go guide: https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
- AWS Docker image on ECR (Elastic Container Registry): https://gallery.ecr.aws/lambda/go
- AWS example of Lambda function written in Go: https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
- Building Go Docker images: https://docs.docker.com/language/golang/build-images/


## Building both Go binary and Docker image

    make


## Manual build steps

### Build Go binary

Create the working directory

    mkdir my-lambda-function

Change to the directory

    cd my-lambda-function

_(Create the files: `main.go` and `Dockerfile`)_

Initialize the module in Go

    go mod init example.com/my-lambda-function

Build the go app for Linux

    GOOS=linux go build

### Build Docker image

To build your Docker image:

    docker build -t my-lambda-function .


## Running the Lambda function

To run your image locally:

    docker run -p 9000:8080 my-lambda-function

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"Tom"}'


## Deploy to AWS Lambda

Parameters used below:

- `[my-lambda-function]`: substitute with your app/repository/image name
- `[AWS Account Number]`: e.g. `123456789012`
- `[Region]`: e.g. `eu-west-1`, `us-east-1`

Create an IAM Policy (e.g. `ECRDockerImageCreation`) with the following permissions:

    ecr:GetAuthorizationToken
    ecr:InitiateLayerUpload
    ecr:CreateRepository

You also need an IAM Role (e.g. `lambda-ex`) for _executing_ the Lambda function.

Build your local Docker image:

    make

### Create a new Lambda function

Authenticate the Docker CLI to your Amazon ECR registry:

    aws ecr get-login-password --region [Region] | docker login --username AWS --password-stdin [AWS Account Number].dkr.ecr.[Region].amazonaws.com

Tag your image to match your repository name:

    docker tag [my-lambda-function]:latest [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

Create an ECR repository:

    aws ecr create-repository --repository-name [my-lambda-function] --region [Region]

Deploy the Docker image to Amazon ECR:

    docker push [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

You should now see your image repository on https://console.aws.amazon.com/ecr/repositories?region=[Region]

    aws lambda create-function --region [Region] --function-name [my-lambda-function] \
      --package-type Image \
      --code ImageUri=[AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function] \
      --role arn:aws:iam::[AWS Account Number]:role/lambda-ex

**Receiving error:**

    An error occurred (InvalidParameterValueException) when calling the CreateFunction operation: Source image [...] is not valid. Provide a valid source image.

### Update an existing Lambda function

Deploy the Docker image to Amazon ECR:

    docker push [AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]:latest

Update the Lambda function:

    aws lambda update-function --region [Region] --function-name [my-lambda-function] \
      --image-uri=[AWS Account Number].dkr.ecr.[Region].amazonaws.com/[my-lambda-function]
