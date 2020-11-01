FROM golang:1.14.0 as builder
WORKDIR /go/src/github.com/takahiro-hayakawa/todo-api-server
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o app main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/takahiro-hayakawa/todo-api-server/app /app
EXPOSE 3000
ENTRYPOINT ["/app"]
