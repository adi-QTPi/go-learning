.DEFAULT_GOAL = build

.PHONY : fmt lint vet clean build
fmt: 
	go fmt ./...
lint:
	golangci-lint run
vet: lint
clean:
	go clean
build: vet # this means that the build target is dependent of the vet target, because we are using 3rd party package for linting, thus it does not depend on it by default, thus the explicit dependency on vet target now. (refer previous Makefile for comparison) 
	go build 