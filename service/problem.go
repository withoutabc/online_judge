package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

func NewProblemServiceImpl() *ProblemDaoImpl {
	return &ProblemDaoImpl{
		ProblemDao: dao.NewProblemDao(),
	}
}

type ProblemDao interface {
	CreateProblem(problem *model.Problem) error
	SearchProblem(request model.ReqSearchProblem) (problems []model.Problem, err error)
	UpdateProblem(problemId int64, problem *model.Problem) error
	DeleteProblem(problemId int64) error
}

type ProblemDaoImpl struct {
	ProblemDao
}

func (p *ProblemDaoImpl) AddProblem(problem model.Problem) int {
	if err := p.ProblemDao.CreateProblem(&problem); err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (p *ProblemDaoImpl) SearchProblem(request model.ReqSearchProblem) (problems []model.Problem, n int) {
	problems, err := p.ProblemDao.SearchProblem(request)
	if len(problems) == 0 {
		return nil, util.NoRecordErrCode
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NoRecordErrCode
		}
		return nil, util.InternalServeErrCode
	}
	return problems, util.NoErrCode
}

func (p *ProblemDaoImpl) UpdateProblem(problemId int64, problem model.Problem) int {
	if err := p.ProblemDao.UpdateProblem(problemId, &problem); err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (p *ProblemDaoImpl) DeleteProblem(problemId int64) int {
	//delete testcase

	//delete problem
	if err := p.ProblemDao.DeleteProblem(problemId); err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}
