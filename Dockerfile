# Compile Go in Dockerfile, see https://docs.aws.amazon.com/lambda/latest/dg/go-image.html

# Use the AWS provided.al2 base image
FROM public.ecr.aws/lambda/provided:al2 as build

# Install the Go compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct

# Copy work files
WORKDIR /app
COPY myFunction/go.mod ./
COPY myFunction/go.sum ./
COPY myFunction/*.go ./

# Compile the Go binary
ENV GOOS=linux
RUN go mod download
ADD . .
RUN go build -o /myFunction

# Copy artifacts to a clean Docker image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /myFunction /myFunction

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "/myFunction" ]
