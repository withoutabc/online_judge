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
	submissions, err := NewJudImpl().SubmissionDao.FindCodeToJudge()
	if err != nil {
		util.Log(err)
		panic(err)
	}
	if len(submissions) == 0 {
		return
	}
	//获取每个提交信息的id
	for _, submission := range submissions {
		var b bool
		_, b = redis.GetValueAndExistence(ctx, submission.SubmissionId)
		//有则直接continue
		if b != false {
			continue
		}
		//没有则写入redis缓存,发送至消息队列
		redis.SetSubmissionId(ctx, submission.SubmissionId)
		ch := ChannelDeclare()
		ExchangeDeclare(ch)
		Publish(ch, submission.Language, submission.SubmissionId)
		log.Println(submission.Language)
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
	SearchSubmissionById(submissionId int64) (submission model.Submission, err error)
	UpdateStatus(submissionId int64, status string) error
}

type TestcaseDao interface {
	SearchTestcase(problemId int64) ([]model.Testcase, error)
}

type InfoDao interface {
	AddCorrect(userId int64) error
	AddScore(userId int64, score int) error
}

type ProblemDao interface {
	AddCorrect(problemId int64) error
	SearchProblem(req model.ReqSearchProblem) (problems []model.Problem, err error)
}

func NewJudImpl() *JudImpl {
	return &JudImpl{
		SubmissionDao: dao.NewSubmissionDao(),
		TestcaseDao:   dao.NewTestDao(),
		ProblemDao:    dao.NewProblemDao(),
		InfoDao:       dao.NewInfoDao(),
	}
}

type JudImpl struct {
	SubmissionDao
	TestcaseDao
	InfoDao
	ProblemDao
}
