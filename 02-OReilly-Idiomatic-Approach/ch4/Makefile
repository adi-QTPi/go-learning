.DEFAULT_GOAL = build

.PHONY : fmt lint vet clean build
fmt: 
	go fmt ./...
lint:
	golangci-lint run
vet: lint
clean:
	go clean
build: vet
	go build 