FROM golang:alpine AS builder
WORKDIR /src/
COPY go.mod .
COPY go.sum .
# RUN go get -d -v
# RUN go mod download
# RUN go test -v -race -timeout 30s ./...
COPY . .
# RUN go mod vendor
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin $GOPATH/src/github.com/omekov/sample
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
