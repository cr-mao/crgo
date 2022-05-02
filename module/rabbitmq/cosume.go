package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var MQConn *amqp.Connection

func init() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "guest", "guest", "127.0.0.1", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	MQConn = conn
	log.Print(MQConn.Major)
}
func GetConn() *amqp.Connection {
	return MQConn
}

func main() {
	conn := GetConn()
	defer conn.Close()
	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	msgs, err := c.Consume("test", "c1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs {

		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
