package dao

import (
	"gorm.io/gorm"
	"online_judge/model"
)

func NewTestDao() *TestDaoImpl {
	return &TestDaoImpl{
		db: DB,
	}
}

type TestDaoImpl struct {
	db *gorm.DB
}

func (t *TestDaoImpl) AddTestcase(testcase *model.Testcase) error {
	result := t.db.Create(&testcase)
	return result.Error
}

func (t *TestDaoImpl) SearchTestcase(problemId int64) ([]model.Testcase, error) {
	var testcases []model.Testcase
	result := t.db.Model(&model.Testcase{}).Where(&model.Testcase{ProblemId: problemId}).Find(&testcases)
	return testcases, result.Error
}

func (t *TestDaoImpl) UpdateTestcase(testcaseId int64, testcase model.Testcase) (int64, error) {
	result := t.db.Model(&model.Testcase{}).Where(&model.Testcase{TestId: testcaseId}).Updates(model.Testcase{
		Input:  testcase.Input,
		Output: testcase.Output,
	})
	return result.RowsAffected, result.Error
}

func (t *TestDaoImpl) DeleteTestcase(tx *gorm.DB, testcaseId int64) (int64, error) {
	result := tx.Model(&model.Testcase{}).Where(&model.Testcase{TestId: testcaseId}).Delete(&model.Testcase{})
	return result.RowsAffected, result.Error
}

func (t *TestDaoImpl) CountTestcase(problemId int64) (int, error) {
	var testcases []model.Testcase
	result := t.db.Model(&model.Testcase{}).Where(&model.Testcase{ProblemId: problemId}).Find(&testcases)
	return len(testcases), result.Error
}
