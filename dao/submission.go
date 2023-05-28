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

func (s *SubmissionDaoImpl) UpdateStatus(submissionId int64, status string) error {
	result := s.db.Model(&model.Submission{}).Where(&model.Submission{SubmissionId: submissionId}).Update("status", status)
	return result.Error
}

func (s *SubmissionDaoImpl) SearchSubmissionById(submissionId int64) (submission model.Submission, err error) {
	result := s.db.Where(&model.Submission{SubmissionId: submissionId}).First(&submission)
	return submission, result.Error
}

func (s *SubmissionDaoImpl) LastSubmission() (int64, error) {
	var Submission model.Submission
	result := s.db.Model(&model.Submission{}).Last(&Submission)
	return 0, result.Error
}

func (s *SubmissionDaoImpl) AddSubmission(tx *gorm.DB, submission *model.Submission) (int64, error) {
	submission.Status = "待测评"
	result := tx.Create(submission)
	if result.Error != nil {
		return 0, result.Error
	}
	//var Submission model.Submission
	//result = s.db.Model(&model.Submission{}).Last(&Submission)
	return 0, nil
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

func (s *SubmissionDaoImpl) FindCodeToJudge() ([]model.Submission, error) {
	var submissions []model.Submission
	result := DB.Model(&model.Submission{}).Where("status", "待测评").Find(&submissions)
	return submissions, result.Error
}
