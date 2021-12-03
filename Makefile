build-zip-file: compile-go
	cd tmp && zip my-lambda-function.zip my-lambda-function && cd ..

compile-go:
	mkdir tmp; \
		cd my-lambda-function; \
		GOOS=linux go build -o ../tmp/

clean-up:
	rm tmp/my-lambda-function
