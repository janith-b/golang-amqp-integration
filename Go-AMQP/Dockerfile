FROM golang:alpine
RUN apk update && apk add curl
WORKDIR /application
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-amqp
ENTRYPOINT ["/go-amqp"]