.DEFAULT_GOAL := build

.PHONY: clean build test vet

vet:
	go vet ./...

test:
	go test -v ./...

build: vet
	go build -o reconsile

clean:
	go clean
