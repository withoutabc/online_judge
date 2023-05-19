package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

type TestDao interface {
	AddTestcase(testcase *model.Testcase) error
	SearchTestcase(problemId int64) ([]model.Testcase, error)
	UpdateTestcase(testcaseId int64, testcase model.Testcase) (int64, error)
	DeleteTestcase(tx *gorm.DB, testcaseId int64) (int64, error)
	CountTestcase(problemId int64) (int, error)
}

func NewTestServiceImpl() *TestDaoImpl {
	return &TestDaoImpl{
		TestDao: dao.NewTestDao(),
		DB:      dao.GetDB(),
	}
}

type TestDaoImpl struct {
	TestDao
	*gorm.DB
}

func (t *TestDaoImpl) AddTestcase(testcase model.Testcase) int {
	err := t.TestDao.AddTestcase(&testcase)
	if err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (t *TestDaoImpl) SearchTestcase(problemId int64) ([]model.Testcase, int) {
	testcases, err := t.TestDao.SearchTestcase(problemId)
	if err != nil {
		return nil, util.InternalServeErrCode
	}
	if len(testcases) == 0 {
		return nil, util.NoRecordErrCode
	}
	return testcases, util.NoErrCode
}

func (t *TestDaoImpl) UpdateTestcase(testcaseId int64, testcase model.Testcase) int {
	count, err := t.TestDao.UpdateTestcase(testcaseId, testcase)
	if err != nil {
		return util.InternalServeErrCode
	}
	if count == 0 {
		return util.UpdateFailErrCode
	}
	return util.NoErrCode
}

func (t *TestDaoImpl) DeleteTestcase(testcaseId int64) int {
	count, err := t.TestDao.DeleteTestcase(t.DB, testcaseId)
	if err != nil {
		return util.InternalServeErrCode
	}
	if count == 0 {
		return util.UpdateFailErrCode
	}
	return util.NoErrCode
}
