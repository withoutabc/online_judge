package judge

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"online_judge/redis"
	"online_judge/util"
	"strconv"
	"time"
)

func Consume(ch *amqp.Channel, language string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	q := QueueDeclare(ch)
	QueueBind(ch, q.Name, language)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			var submissionId int64
			submissionId, err = strconv.ParseInt(string(d.Body), 10, 64)
			if err != nil {
				util.Log(err)
				continue
			}
			//根据id查询submission的code
			//var code string
			//code, err = NewSubmissionServiceImpl().SubmissionDao.SearchSubmissionById(submissionId)
			//if err != nil {
			//	//没查询到跳过
			//	util.Log(err)
			//	continue
			//}
			//放到容器里面跑，得到输出结果

			//比对结果，修改测评状态

			//修改测评结果

			//删除redis缓存
			redis.DeleteSubmissionId(ctx, submissionId)
			//确认
			d.Ack(false)
		}
	}()

	<-forever
}
