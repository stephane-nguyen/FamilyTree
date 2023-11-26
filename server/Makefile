BINARY_NAME=server
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 G00S=linux go build -o target/$(BINARY_NAME)-darwin main.go
	GOARCH=amd64 G00S=linux go build -o target/$(BINARY_NAME)-linux main.go
	GOARCH=amd64 G00S=linux go build -o target/$(BINARY_NAME)-windows main.go

run: build
	./target/$(BINARY_NAME)-darwin

clean:
	go clean
	./target/$(BINARY_NAME)-darwin
	./target/$(BINARY_NAME)-linux
	./target/$(BINARY_NAME)-windows