FROM golang:alpine AS builder
LABEL Omekov Azamat <umekovazamat@gmail.com>
WORKDIR $GOPATH/src/github.com/omekov/sample/cmd
COPY go.mod .
COPY go.sum .
# RUN go get -d -v
RUN go mod download
# RUN go test -v -race -timeout 30s ./...
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app $GOPATH/src/github.com/omekov/sample/cmd/

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]