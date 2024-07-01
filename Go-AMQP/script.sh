#!/bin/bash
export API_ENDPOINT=0.0.0.0:8000
export BASE_PATH=/home/janith/go-gin-rabbitmq-api
export AMQP_URL=amqp://192.168.1.171:5672
export AMQP_EXCHANGE_NAME=exchange_upload_success
export AMQP_EXCHANGE_KIND=fanout
export AMQP_QUEUE_NAME=queue_upload_success
export AMQP_ROUTING_KEY=upload_success

go run . $1