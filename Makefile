.PHONY: build

default: run

test: unit-tests

unit-tests:
	go fmt ./...
	go test -vet all -shuffle=on ./...

build: test
	go build -o wordle ./cmd/wordle

mod:
	go mod vendor -v

tidy:
	go mod tidy -v

run: build
	@echo "-------------------"
	@./wordle
