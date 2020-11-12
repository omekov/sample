FROM golang:alpine AS builder
LABEL Omekov Azamat <umekovazamat@gmail.com>
WORKDIR $GOPATH/src/github.com/omekov/sample
COPY go.mod .
COPY go.sum .
# RUN go get -d -v
# RUN go mod download
# RUN go test -v -race -timeout 30s ./...
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/golang $GOPATH/src/github.com/omekov/sample/cmd/auth

FROM scratch
COPY --from=builder /go/bin/golang /go/bin/golang
ENTRYPOINT ["/go/bin/golang"]