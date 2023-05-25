package judge

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"online_judge/logs"
	"strconv"
)

func Publish(ch *amqp.Channel, language string, submissionId int64) {
	if err := ch.PublishWithContext(context.Background(),
		"language", // exchange
		language,   // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(strconv.FormatInt(submissionId, 10)),
		}); err != nil {
		msg := fmt.Sprintf("publish %s error", strconv.FormatInt(submissionId, 10))
		logs.Log().Error(msg, zap.Error(err))
		panic(err)
	}
}
