FROM public.ecr.aws/lambda/go:1

# Copy function to /var/task
COPY myFunction ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "myFunction" ]
