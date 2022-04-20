.PHONY: clean lint vet fmt build test docker

all: clean lint vet fmt build test

clean:
	rm -rf ./bin

lint:
	golangci-lint run

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | xargs -L1 gofmt -w

build: build-darwin build-linux

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/tks-darwin-amd64 ./cmd/server/
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/tks-appclient-darwin-amd64 ./examples/client.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tks-linux-amd64 ./cmd/server/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tks-appclient-linux-amd64 ./examples/client.go

test:
	go test -v ./...

docker:
	docker build --no-cache -t tks-info/tks-info -f Dockerfile .
