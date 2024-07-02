package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type Log struct {
	Kind       string   `json:"kind"`
	Apiversion string   `json:"apiVersion"`
	Level      string   `json:"level"`
	AuditID    string   `json:"auditID"`
	Stage      string   `json:"stage"`
	RequestURI string   `json:"requestURI"`
	Verb       string   `json:"verb"`
	User       User     `json:"user"`
	SourceIPs  []string `json:"sourceIPs"`

	UserAgent string    `json:"userAgent"`
	ObjectRef ObjectRef `json:"objectRef"`

	ResponseStatus ResponseStatus `json:"responseStatus"`

	RequestReceivedTimestamp string                 `json:"requestReceivedTimestamp"`
	StageTimestamp           string                 `json:"stageTimestamp"`
	Annotations              map[string]interface{} `json:"annotations"`
}

type User struct {
	UserName string   `json:"username"`
	Groups   []string `json:"groups"`
}

type ObjectRef struct {
	Resource   string `json:"resource"`
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
	ApiVersion string `json:"apiVersion"`
}

type ResponseStatus struct {
	Metadata map[string]interface{} `json:"metadata"`
	Code     int                    `json:"code"`
}

func readLogFile(message string) {
	msg := Message{}
	logLines := Log{}
	json.Unmarshal([]byte(message), &msg)
	f, err := os.Open(msg.FullFilePath)
	if err != nil {
		log.Println(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		e := json.Unmarshal([]byte(s.Text()), &logLines)
		if e != nil {
			log.Println(e)
			break
		}
		fmt.Println(logLines.RequestReceivedTimestamp, logLines.RequestURI)
	}

}

func (r RabbitMQ_Properties) consumeQueue(ch *amqp091.Channel) {
	msgs, _ := ch.Consume(
		r.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	for d := range msgs {
		log.Println("MESSAGE : ", string(d.Body))
		readLogFile(string(d.Body))
		d.Ack(false)
	}

}
