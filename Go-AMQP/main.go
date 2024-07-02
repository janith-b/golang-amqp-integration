package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

var API_BIND_IP string = os.Getenv("API_BIND_IP")
var API_BIND_PORT string = os.Getenv("API_BIND_PORT")
var BASE_PATH string = os.Getenv("BASE_PATH")
var AMQP_URL string = os.Getenv("AMQP_URL")
var AMQP_EXCHANGE_NAME string = os.Getenv("AMQP_EXCHANGE_NAME")
var AMQP_EXCHANGE_KIND string = os.Getenv("AMQP_EXCHANGE_KIND")
var AMQP_QUEUE_NAME string = os.Getenv("AMQP_QUEUE_NAME")
var AMQP_ROUTING_KEY string = os.Getenv("AMQP_ROUTING_KEY")

var rabbitmqProperties RabbitMQ_Properties = RabbitMQ_Properties{
	Url:          AMQP_URL,
	ExchName:     AMQP_EXCHANGE_NAME,
	ExchKind:     AMQP_EXCHANGE_KIND,
	ExchDurable:  false,
	QueueName:    AMQP_QUEUE_NAME,
	QueueDurable: false,
	RoutingKey:   AMQP_ROUTING_KEY,
}

func main() {
	ch := rabbitmqProperties.initRabbitMQConnection()
	if os.Args[1] == "server" {
		router := gin.Default()
		router.GET("/health", healthHandler())
		router.POST("/upload", uploadFileHander(BASE_PATH, ch))
		router.Run(API_BIND_IP + ":" + API_BIND_PORT)
	} else if os.Args[1] == "consumer" {
		log.Println("Started consuming messages")
		rabbitmqProperties.consumeQueue(ch)
	}
}

func healthHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(200, "Application is Healthy")
	}
}

func uploadFileHander(basePath string, ch *amqp091.Channel) func(c *gin.Context) {
	return func(c *gin.Context) {

		f, e1 := c.FormFile("file")
		if e1 != nil {
			log.Println(e1)
		}
		e2 := c.SaveUploadedFile(f, basePath+"/"+f.Filename)
		if e2 != nil {
			log.Println(e2)
		} else {
			c.String(201, "File uploaded successfully")
			message := Message{
				FileName:     f.Filename,
				FullFilePath: basePath + "/" + f.Filename,
				FileSize:     f.Size,
				Timestamp:    time.Now().Format("2006-01-02T15:04:05"),
			}
			bs, _ := json.Marshal(message)
			rabbitmqProperties.writeToQueue(ch, bs)

		}
	}
}
