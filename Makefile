test:
	go test ./...

build: 
	go mod download && go build -o bin/cmd ./cmd

run:
	go run ./cmd

lint:
	golangci-lint run
