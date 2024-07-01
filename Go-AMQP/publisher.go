package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ_Properties struct {
	Url          string
	ExchName     string
	ExchKind     string
	ExchDurable  bool
	QueueName    string
	QueueDurable bool
	RoutingKey   string
}

type Message struct {
	FileName     string `json:"fileName"`
	FullFilePath string `json:"fullFilePath"`
	Timestamp    string `json:"timeStamp"`
	FileSize     int64  `json:"fileSize"`
}

func (r RabbitMQ_Properties) initRabbitMQConnection() *amqp091.Channel {
	conn, e1 := amqp091.Dial(r.Url)
	if e1 != nil {
		log.Println(e1)
	}
	ch, e2 := conn.Channel()
	if e2 != nil {
		log.Println(e2)
	}
	e3 := ch.ExchangeDeclare(
		r.ExchName,
		r.ExchKind,
		r.ExchDurable,
		false,
		false,
		false,
		nil,
	)
	if e3 != nil {
		log.Println(e3)
	}
	que, e4 := ch.QueueDeclare(
		r.QueueName,
		r.QueueDurable,
		false,
		false,
		false,
		nil,
	)

	if e4 != nil {
		log.Println(e4)
	}

	ch.QueueBind(
		que.Name,
		r.RoutingKey,
		r.ExchName,
		false,
		nil,
	)

	return ch
}

func (r RabbitMQ_Properties) writeToQueue(ch *amqp091.Channel, message []byte) {
	ch.Publish(r.ExchName,
		r.RoutingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
}
