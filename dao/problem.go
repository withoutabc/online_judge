package dao

import (
	"gorm.io/gorm"
	"online_judge/model"
	"time"
)

func NewProblemDao() *ProblemDaoImpl {
	return &ProblemDaoImpl{
		db: DB,
	}
}

type ProblemDaoImpl struct {
	db *gorm.DB
}

func (p *ProblemDaoImpl) AddCorrect(problemId int64) error {
	result := p.db.Model(&model.Problem{}).Where(&model.Problem{ProblemId: problemId}).UpdateColumn("correct", gorm.Expr("correct + ?", 1))
	return result.Error
}

func (p *ProblemDaoImpl) SearchTitleExist(title string) (bool, error) {
	var problem model.Problem
	result := p.db.Model(&model.Problem{}).Where(&model.Problem{Title: title}).First(&problem)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return false, result.Error
		}
		return false, nil
	}
	return true, nil
}

func (p *ProblemDaoImpl) SearchExistById(problemId int64) (bool, error) {
	var problem model.Problem
	result := p.db.Model(&model.Problem{}).Where(&model.Problem{ProblemId: problemId}).First(&problem)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return false, result.Error
		}
		return false, nil
	}
	return true, nil
}

func (p *ProblemDaoImpl) AddProblemSubmit(tx *gorm.DB, problemId int64) error {
	result := tx.Model(&model.Problem{}).Where(&model.Problem{ProblemId: problemId}).UpdateColumn("submit", gorm.Expr("submit + ?", 1))
	return result.Error
}

func (p *ProblemDaoImpl) CreateProblem(problem *model.Problem) error {
	result := DB.Create(&problem)
	return result.Error
}

func (p *ProblemDaoImpl) SearchProblem(req model.ReqSearchProblem) (problems []model.Problem, err error) {
	cond := model.Problem{}
	if req.UserId != 0 {
		cond.UserId = req.UserId
	}
	if req.ProblemId != 0 {
		cond.ProblemId = req.ProblemId
	}
	if req.Level != "" {
		cond.Level = req.Level
	}
	if req.Keyword != "" {
		result := p.db.Where("title LIKE ?", "%"+req.Keyword+"%").Find(&problems)
		return problems, result.Error
	}
	result := p.db.Where(&cond).Find(&problems)
	return problems, result.Error
}

func (p *ProblemDaoImpl) UpdateProblem(problemId int64, problem model.Problem) (int64, error) {
	result := p.db.Where("problem_id", problemId).Updates(&model.Problem{
		Title:             problem.Title,
		Description:       problem.Description,
		DescriptionInput:  problem.DescriptionInput,
		DescriptionOutput: problem.DescriptionOutput,
		SampleInput:       problem.SampleInput,
		SampleOutput:      problem.SampleOutput,
		Level:             problem.Level,
		UpdateTime:        time.Now().Format("2006-01-02 15:04:05"),
	})
	return result.RowsAffected, result.Error
}

func (p *ProblemDaoImpl) DeleteProblem(tx *gorm.DB, problemId int64) (int64, error) {
	result := tx.Model(&model.Problem{}).Where(&model.Problem{ProblemId: problemId}).Delete(&model.Problem{})
	return result.RowsAffected, result.Error
}
