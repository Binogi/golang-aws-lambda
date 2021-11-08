FROM public.ecr.aws/lambda/go:latest

# Copy function to /var/task
COPY src/myFunction.go ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "myFunction.go" ]
