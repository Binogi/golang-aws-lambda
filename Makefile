build-docker-image: compile-go
	docker build -t my-lambda-function .

compile-go:
	mkdir tmp; \
		cd my-lambda-function; \
		GOOS=linux go build -o ../tmp/

clean-up:
	rm tmp/my-lambda-function
