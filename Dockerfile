FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/omekov/sample

# COPY go.mod .
# COPY go.sum .
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app $GOPATH/src/github.com/omekov/sample/cmd/sample

# RUN go test -v -race -timeout 30s ./...

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]