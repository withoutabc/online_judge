package dao

import (
	"gorm.io/gorm"
	"online_judge/model"
)

func NewSubmissionDao() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		db: DB,
	}
}

type SubmissionDaoImpl struct {
	db *gorm.DB
}

func (s *SubmissionDaoImpl) AddSubmission(tx *gorm.DB, submission *model.Submission) error {
	submission.Status = "待测评"
	result := tx.Create(submission)
	return result.Error
}

func (s *SubmissionDaoImpl) SearchSubmission(req model.ReqSearchSubmission) (submissions []model.Submission, err error) {
	cond := model.Submission{UserId: req.UserId}
	if req.ProblemId != 0 {
		cond.ProblemId = req.ProblemId
	}
	if req.Language != "" {
		cond.Language = req.Language
	}
	if req.Status != "" {
		cond.Status = req.Status
	}
	if req.From != "" && req.To != "" {
		fromStr := req.From
		toStr := req.To
		cond.SubmitTime = "update_time BETWEEN '" + fromStr + "' AND '" + toStr + "'"
	} else if req.From != "" {
		fromStr := req.From
		cond.SubmitTime = "update_time >= '" + fromStr + "'"
	} else if req.To != "" {
		toStr := req.To
		cond.SubmitTime = "update_time <= '" + toStr + "'"
	}
	result := s.db.Model(&model.Submission{}).Where(&cond).Find(&submissions)
	return submissions, result.Error
}
