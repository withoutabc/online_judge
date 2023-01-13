package service

import (
	"online_judge/dao"
	"online_judge/model"
)

func Submit(s model.Submission) (err error) {
	err = dao.InsertSubmission(s)
	return
}

func ViewResult(s model.Submission) ([]model.Submission, error) {
	submissions, err := dao.QueryResult(s)
	return submissions, err
}
