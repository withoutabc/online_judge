package main

import (
	"context"
	"online_judge/dao"
	"online_judge/exert/judge"
	"online_judge/model"
	"online_judge/redis"
	"online_judge/util"
	"time"
)

func main() {
	dao.InitDB()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
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
				value, b, err = redis.GetValueAndExistence(ctx, submission.SubmissionId)
				if err != nil {
					util.Log(err)
					panic(err)
				}
				//有则发送至消息队列，然后直接continue
				if b == true || value == "true" {
					continue
				}
				//没有则写入redis缓存,发送至消息队列
				err = redis.SetSubmissionId(ctx, submission.SubmissionId)
				if err != nil {
					util.Log(err)
					panic(err)
				}
				ch := judge.ChannelDeclare()
				judge.ExchangeDeclare(ch)
				judge.Publish(ch, submission.Language, submission.SubmissionId)
			}
		}
	}

}

type SubmissionDao interface {
	FindCodeToJudge() ([]model.Submission, error)
}

func NewSubmissionServiceImpl() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		SubmissionDao: dao.NewSubmissionDao(),
	}
}

type SubmissionDaoImpl struct {
	SubmissionDao
}
