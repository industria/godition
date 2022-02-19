.DEFAULT_GOAL := build

.PHONY: clean build test vet run-decks

vet:
	go vet ./...

test:
	go test -v ./...

build: vet
	go build -o reconsile

clean:
	go clean

run-decks:
	go run ./cmd/decks/