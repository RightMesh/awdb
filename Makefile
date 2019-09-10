all: build

build: lint
	go build -v .

clean:
	go clean

lint:
	${GOPATH}/bin/golint ./...

run:
	go run cmd/main.go
