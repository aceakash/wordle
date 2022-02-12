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

dist: test
	env GOOS=darwin GOARCH=amd64 go build -o dist/wordle-darwin-amd64 ./cmd/wordle
	env GOOS=darwin GOARCH=arm64 go build -o dist/wordle-darwin-arm64 ./cmd/wordle
	env GOOS=linux GOARCH=amd64 go build -o dist/wordle-linux-amd64 ./cmd/wordle
	env GOOS=windows GOARCH=amd64 go build -o dist/wordle-windows-amd64 ./cmd/wordle

