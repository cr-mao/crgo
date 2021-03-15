package rabbitmq

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"crgo/infra/log"
)

const (
	RABBITMQ = "rabbitmq"
	DSN      = "dsn"
)

type AMQPMap struct {
	mapping map[string]*amqp.Connection
}

var dsnMap map[string]string
var amqpMap AMQPMap
var once sync.Once

func Init() {
	once.Do(func() {
		mapping := make(map[string]*amqp.Connection)
		dsnMap = make(map[string]string)

		for bind := range viper.GetStringMap(RABBITMQ) {
			dsn := viper.GetString(strings.Join([]string{RABBITMQ, bind, DSN}, "."))
			conn, err := amqp.Dial(dsn)
			if err != nil {
				panic(err)
			}
			log.Debugf("preparing RabbitMQ amqp.Connection -> %s @ %s", bind, dsn)
			mapping[bind] = conn
			dsnMap[bind] = dsn
		}
		amqpMap = AMQPMap{mapping: mapping}
	})
}
