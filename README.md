# Go functions on AWS Lambda in Docker

Based on [aws-lambda-go on Docker Hub](https://hub.docker.com/r/amazon/aws-lambda-go) and [this example Gist](https://gist.github.com/josephspurrier/05b9126279703a81122cba198df50d6f).

See also:

- AWS Go guide: https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
- Docker image: https://gallery.ecr.aws/lambda/go
- AWS Lambda function handler in Go: https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html

## Build Go binary

Create the working directory

    mkdir myFunction

Change to the directory

    cd myFunction

_(Create the files: `main.go` and `Dockerfile`)_

Initialize the module in Go

    go mod init example.com/myFunction

Build the go app for Linux

    GOOS=linux go build

## Build Docker image

To build your Docker image:

    docker build -t my-lambda-app .

To run your image locally:

    docker run -p 9000:8080 my-lambda-app

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"payload":"hello world!"}'
