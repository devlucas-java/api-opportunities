
.PHONY: default run build test docs clean
# Variables
APP_NAME=api-opportunities
VERSION=0.1.0

default: run

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test -v ./...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
