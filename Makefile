.DEFAULT_GOAL := build

.PHONY: clean build test vet run-decks

vet:
	go vet ./...

test:
	go test -v ./...

build: vet
	go build -o decks -v ./cmd/decks 

clean:
	go clean
	rm -f decks

run-decks:
	go run ./cmd/decks/