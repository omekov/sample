.PHONY: run
run:
				swag init
				go run *.go -env
.PHONY: build
build:
				go build -v 
.PHONY: test
test:
				go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build