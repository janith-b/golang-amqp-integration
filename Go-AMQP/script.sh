#!/bin/bash
export API_ENDPOINT=0.0.0.0:8000
export BASE_PATH=/home/janith/go-gin-rabbitmq-api
export AMQP_URL=amqp://192.168.1.172:5672
export AMQP_EXCHANGE_NAME=exchange_upload_success
export AMQP_EXCHANGE_KIND=fanout
export AMQP_QUEUE_NAME=queue_upload_success
export AMQP_ROUTING_KEY=upload_success
export API_BIND_IP=0.0.0.0
export API_BIND_PORT=8000

# go run . $1


docker run \
-e API_BIND_IP=0.0.0.0 \
-e API_BIND_PORT=8000 \
-e BASE_PATH=/base \
-e AMQP_URL=amqp://192.168.1.172:5672 \
-e AMQP_EXCHANGE_NAME=exchange_upload_success \
-e AMQP_EXCHANGE_KIND=fanout \
-e AMQP_QUEUE_NAME=queue_upload_success \
-e AMQP_ROUTING_KEY=upload_success \
-p 8000:8000 \
-v ./base-path:/base \
go-amqp server