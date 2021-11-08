# Go functions on AWS Lambda in Docker

- AWS Go guide: https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
- aws-lambda-go on Docker Hub: https://hub.docker.com/r/amazon/aws-lambda-go
- Docker image: https://gallery.ecr.aws/lambda/go
- AWS Lambda function handler in Go: https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
- Example Gist: https://gist.github.com/josephspurrier/05b9126279703a81122cba198df50d6f

To build your image:

    docker build -t my-lambda-app .

To run your image locally:

    docker run -p 9000:8080 my-lambda-app

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"payload":"hello world!"}'
