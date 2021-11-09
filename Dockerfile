# Use the AWS Go Docker base image
FROM public.ecr.aws/lambda/go:1

# Copy the compiled Go binary file to the Lambda task folder (/var/task)
COPY tmp/my-lambda-function ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "my-lambda-function" ]
