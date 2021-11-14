.PHONY: run
run:
				swag init -g cmd/sample/*.go
				go run ./cmd/sample/*.go
.PHONY: build
build:
				go build -o sample ./cmd/sample
.PHONY: test
test:
				go test -v -race -timeout 30s ./...
.PHONY: docker
docker:
				cd deployments && docker-compose up
.DEFAULT_GOAL := build