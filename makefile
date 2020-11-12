.PHONY: run
run:
				swag init
				go run ./cmd/auth/*.go -env
.PHONY: build
build:
				go build ./cmd/auth/ 
.PHONY: test
test:
				go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build