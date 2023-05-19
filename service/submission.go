package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

type SubmissionDao interface {
	AddSubmission(submission *model.Submission) error
	SearchSubmission(req model.ReqSearchSubmission) (submissions []model.Submission, err error)
}

func NewSubmissionServiceImpl() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		SubmissionDao: dao.NewSubmissionDao(),
	}
}

type SubmissionDaoImpl struct {
	SubmissionDao
}

func (s *SubmissionDaoImpl) AddSubmission(submission model.Submission) int {
	if err := s.SubmissionDao.AddSubmission(&submission); err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (s *SubmissionDaoImpl) SearchSubmission(req model.ReqSearchSubmission) ([]model.Submission, int) {
	submissions, err := s.SubmissionDao.SearchSubmission(req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NoRecordErrCode
		}
		return nil, util.InternalServeErrCode
	}
	if len(submissions) == 0 {
		return nil, util.NoRecordErrCode
	}
	return submissions, util.NoErrCode
}
