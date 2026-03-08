.PHONY: serve worker build lint test tidy fmt

export GOFLAGS := -mod=mod

serve:
	go run ./cmd/serve

worker:
	go run ./cmd/worker

build:
	mkdir -p bin
	go build -o bin/serve ./cmd/serve
	go build -o bin/worker ./cmd/worker

lint:
	golangci-lint run ./...

test:
	go test ./... -v -race -cover

tidy:
	go mod tidy

fmt:
	gofmt -w .

.env:
	cp .env.example .env
