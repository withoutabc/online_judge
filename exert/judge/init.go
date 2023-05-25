package judge

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"online_judge/logs"
)

var Conn *amqp.Connection

func InitRabbitMq() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logs.Log().Error("fail to init rabbitMq", zap.Error(err))
		panic(err)
	}
	Conn = conn
}

func ChannelDeclare() (ch *amqp.Channel) {
	ch, err := Conn.Channel()
	if err != nil {
		logs.Log().Error("fail to open a channel", zap.Error(err))
		panic(err)
	}
	return ch
}

func QueueDeclare(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logs.Log().Error("fail to declare a channel", zap.Error(err))
		panic(err)
	}
	return q
}

func ExchangeDeclare(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		"language", // name
		"direct",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		logs.Log().Error("fail to declare a channel", zap.Error(err))
		panic(err)
	}
}

func QueueBind(ch *amqp.Channel, queueName string, language string) {
	err := ch.QueueBind(
		queueName,  // queue name
		language,   // routing key
		"language", // exchange
		false,
		nil,
	)
	if err != nil {
		logs.Log().Error("fail to bind queue", zap.Error(err))
	}
}
