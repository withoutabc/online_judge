package dao

import (
	"fmt"
	"gorm.io/gorm"
	"online_judge/model"
)

func NewProblemDao() *ProblemDaoImpl {
	return &ProblemDaoImpl{
		db: DB,
	}
}

type ProblemDaoImpl struct {
	db *gorm.DB
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
	if req.From != "" && req.To != "" {
		fromStr := req.From
		toStr := req.To
		cond.UpdateTime = "update_time BETWEEN '" + fromStr + "' AND '" + toStr + "'"
	} else if req.From != "" {
		fromStr := req.From
		cond.UpdateTime = "update_time >= '" + fromStr + "'"
	} else if req.To != "" {
		toStr := req.To
		cond.UpdateTime = "update_time <= '" + toStr + "'"
	}
	if req.Keyword != "" {
		cond.Title = fmt.Sprintf("%%%s%%", req.Keyword)
	}
	result := p.db.Model(&model.Problem{}).Where(&cond).Find(&problems)
	return problems, result.Error
}

func (p *ProblemDaoImpl) UpdateProblem(problemId int64, problem *model.Problem) error {
	result := p.db.Take(&model.Problem{}).Where(&model.Problem{ProblemId: problemId}).Updates(model.Problem{
		Title:             problem.Title,
		Description:       problem.Description,
		DescriptionInput:  problem.DescriptionInput,
		DescriptionOutput: problem.DescriptionOutput,
		SampleInput:       problem.SampleInput,
		SampleOutput:      problem.SampleOutput,
		Level:             problem.Level,
	})
	return result.Error
}

func (p *ProblemDaoImpl) DeleteProblem(problemId int64) error {
	result := p.db.Table("problems").Delete(&model.Problem{ProblemId: problemId})
	return result.Error
}
