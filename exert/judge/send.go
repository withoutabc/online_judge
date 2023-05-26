package judge

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/redis"
	"online_judge/util"
	"strconv"
	"time"
)

func Produce() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//读取待测评提交信息，
	submissions, err := NewSubmissionServiceImpl().SubmissionDao.FindCodeToJudge()
	if err != nil {
		util.Log(err)
		panic(err)
	}
	//获取每个提交信息的id
	for _, submission := range submissions {
		var (
			value string
			b     bool
		)
		value, b = redis.GetValueAndExistence(ctx, submission.SubmissionId)
		//有则直接continue
		if b == true || value == "true" {
			continue
		}
		//没有则写入redis缓存,发送至消息队列
		redis.SetSubmissionId(ctx, submission.SubmissionId)
		log.Println(submission.SubmissionId)
		ch := ChannelDeclare()
		ExchangeDeclare(ch)
		Publish(ch, submission.Language, submission.SubmissionId)
	}
}

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
		util.Log(err)
		panic(err)
	}
}

type SubmissionDao interface {
	FindCodeToJudge() ([]model.Submission, error)
	SearchSubmissionById(submissionId int64) (code string, err error)
}

func NewSubmissionServiceImpl() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		SubmissionDao: dao.NewSubmissionDao(),
	}
}

type SubmissionDaoImpl struct {
	SubmissionDao
}
