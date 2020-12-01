.PHONY: run
run:
				swag init -g cmd/main.go
				go run ./cmd/main.go
.PHONY: build
build:
				go build ./cmd 
.PHONY: test
test:
				go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build