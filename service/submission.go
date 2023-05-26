package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

type SubmissionDao interface {
	AddSubmission(tx *gorm.DB, submission *model.Submission) (int64, error)
	SearchSubmission(req model.ReqSearchSubmission) (submissions []model.Submission, err error)
	LastSubmission() (int64, error)
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
	_, err := s.SubmissionDao.AddSubmission(tx, &submission)
	if err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if err = s.ProblemDao.AddProblemSubmit(tx, submission.ProblemId); err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	_, err = s.SubmissionDao.LastSubmission()
	if err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (s *SubmissionDaoImpl) SearchSubmission(req model.ReqSearchSubmission) ([]model.SubAndPro, int) {
	var subAndPro []model.SubAndPro
	submissions, err := s.SubmissionDao.SearchSubmission(req)
	if err != nil {
		return nil, util.InternalServeErrCode
	}
	if len(submissions) == 0 {
		return nil, util.NoRecordErrCode
	}
	subAndPro = make([]model.SubAndPro, len(submissions))
	//对于每个submission找problem
	for i, submission := range submissions {
		problems, err := s.ProblemDao.SearchProblem(model.ReqSearchProblem{ProblemId: submission.ProblemId})
		if err != nil {
			return nil, util.InternalServeErrCode
		}
		subAndPro[i].Submission = submission
		if len(problems) == 0 {
			continue
		}
		subAndPro[i].Problem = problems[0]
	}

	return subAndPro, util.NoErrCode
}
