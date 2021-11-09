build-docker-image: compile-go clean-up
	docker build -t my-lambda-app .

compile-go:
	cd myFunction; \
		GOOS=linux go build

clean-up:
	rm myFunction/myFunction
