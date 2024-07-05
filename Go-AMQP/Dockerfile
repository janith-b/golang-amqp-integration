FROM golang:alpine AS builder
WORKDIR /application
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-amqp

FROM busybox
COPY --from=builder /go-amqp /go-amqp
ENTRYPOINT ["/go-amqp"]