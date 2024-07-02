# Golang-AMQP Integration
## Introduction

A message queue is a system that enables asynchronous communication between services, commonly used in serverless and microservices architectures. It decouples the services allowing them to to scale independently while increasing the reliability of the system. In this project, I developed a simple log parser in Golang to demonstrate the potential of message queuing and event-driven architecture.

This system consists of three major components:
- RabbitMQ instance
- REST API along with message the Producer
- Message Consumer and Log parser

All the components mentioned above are containerized and can be easily recreated on any system using Docker Compose.

## Design and Working Priciple
The REST API accepts Kubernetes Audit logs and stores them in persistent storage. After a log file is uploaded, the API publishes a message to a queue in a RabbitMQ instance. This message is then consumed by a Consumer, which also has access to the persistent storage. The Consumer retrieves the log file, parses it, and prints the Request URIs from the specific Kubernetes Audit log file.
#### Sample Message :
```
{"fileName":"audit-2024-06-28T04-57-43.125.log","fullFilePath":"/application/log-files/audit-2024-06-28T04-57-43.125.log","timeStamp":"2024-07-02T12:11:08","fileSize":1705681}
```

## Commands
#### Spinup the Containers
```
docker compose up -d
```
#### Stop the Containers
```
docker compose down -d
```