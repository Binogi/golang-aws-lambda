build-docker-image: compile-go
	docker build -t my-lambda-function .

compile-go:
	cd my-lambda-function; \
		GOOS=linux go build

clean-up:
	rm my-lambda-function/my-lambda-function
