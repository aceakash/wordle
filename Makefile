.PHONY: build

default: run

tests: unit-tests

unit-tests:
	go fmt ./...
	go test -vet all -shuffle=on ./...

build: tests
	go build -o wordle ./cmd/wordle

mod:
	go mod vendor -v

tidy:
	go mod tidy -v

run: build
	@echo "-------------------"
	@./wordle
