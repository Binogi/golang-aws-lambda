# Go functions on AWS Lambda in Docker

- https://hub.docker.com/r/amazon/aws-lambda-go
- https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html

To build your image:

    docker build -t my-lambda-app .

To run your image locally:

    docker run -p 9000:8080 my-lambda-app

In a separate terminal, you can then locally invoke the function using cURL:

    curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"payload":"hello world!"}'
