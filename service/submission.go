package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

type SubmissionDao interface {
	AddSubmission(tx *gorm.DB, submission *model.Submission) error
	SearchSubmission(req model.ReqSearchSubmission) (submissions []model.Submission, err error)
}

func NewSubmissionServiceImpl() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		SubmissionDao: dao.NewSubmissionDao(),
		ProblemDao:    dao.NewProblemDao(),
		DB:            dao.GetDB(),
	}
}

type SubmissionDaoImpl struct {
	SubmissionDao
	ProblemDao
	*gorm.DB
}

func (s *SubmissionDaoImpl) AddSubmission(submission model.Submission) int {
	tx := s.DB.Begin()
	if err := s.SubmissionDao.AddSubmission(tx, &submission); err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if err := s.ProblemDao.AddProblemSubmit(tx, submission.ProblemId); err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (s *SubmissionDaoImpl) SearchSubmission(req model.ReqSearchSubmission) ([]model.Submission, int) {
	submissions, err := s.SubmissionDao.SearchSubmission(req)
	if err != nil {
		return nil, util.InternalServeErrCode
	}
	if len(submissions) == 0 {
		return nil, util.NoRecordErrCode
	}
	return submissions, util.NoErrCode
}
