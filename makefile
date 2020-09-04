.PHONY: run
run:
				swag init
.PHONY: build
build:
				go build -v
.PHONY: test
test:
				go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build