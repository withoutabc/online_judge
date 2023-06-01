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
		TestDao:    dao.NewTestDao(),
		DB:         dao.GetDB(),
	}
}

type ProblemDao interface {
	CreateProblem(problem *model.Problem) error
	SearchTitleExist(title string) (bool, error)
	SearchProblem(request model.ReqSearchProblem) (problems []model.Problem, err error)
	SearchExistById(problemId int64) (bool, error)
	UpdateProblem(problemId int64, problem model.Problem) (int64, error)
	DeleteProblem(tx *gorm.DB, problemId int64) (int64, error)
	AddProblemSubmit(tx *gorm.DB, problemId int64) error
}

type ProblemDaoImpl struct {
	ProblemDao
	TestDao
	*gorm.DB
}

func (p *ProblemDaoImpl) AddProblem(problem model.Problem) int {
	b, err := p.ProblemDao.SearchTitleExist(problem.Title)
	if err != nil {
		return util.InternalServeErrCode
	}
	if b == true {
		return util.RepeatedTitleErrCode
	}
	if err = p.ProblemDao.CreateProblem(&problem); err != nil {
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
		return nil, util.InternalServeErrCode
	}
	return problems, util.NoErrCode
}

func (p *ProblemDaoImpl) UpdateProblem(problemId int64, problem model.Problem) int {
	b, err := p.ProblemDao.SearchExistById(problemId)
	if err != nil {
		return util.InternalServeErrCode
	}
	if b != true {
		return util.NoRecordErrCode
	}
	b, err = p.ProblemDao.SearchTitleExist(problem.Title)
	if err != nil {
		return util.InternalServeErrCode
	}
	if b == true {
		return util.RepeatedTitleErrCode
	}
	count, err := p.ProblemDao.UpdateProblem(problemId, problem)
	if err != nil {
		return util.InternalServeErrCode
	}
	if count == 0 {
		return util.UpdateFailErrCode
	}
	return util.NoErrCode
}

func (p *ProblemDaoImpl) DeleteProblem(problemId int64) int {
	b, err := p.ProblemDao.SearchExistById(problemId)
	if err != nil {
		return util.InternalServeErrCode
	}
	if b != true {
		return util.NoRecordErrCode
	}
	tx := p.DB.Begin()
	//delete testcase
	testcases, err := p.TestDao.SearchTestcase(problemId)
	if err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if len(testcases) != 0 {
		for _, testcase := range testcases {
			count, err := p.TestDao.DeleteTestcase(tx, testcase.TestId)
			if err != nil {
				tx.Rollback()
				return util.InternalServeErrCode
			}
			if count == 0 {
				tx.Rollback()
				return util.UpdateFailErrCode
			}
		}
	}
	//delete problem
	count, err := p.ProblemDao.DeleteProblem(tx, problemId)
	if err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	if count == 0 {
		tx.Rollback()
		return util.UpdateFailErrCode
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}
