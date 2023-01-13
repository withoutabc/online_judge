package service

import (
	"online_judge/dao"
	"online_judge/model"
)

func AddProblem(p model.Problem) (err error) {
	err = dao.InsertProblem(p)
	return
}

func ViewProblems() (problems []model.Problem, err error) {
	problems, err = dao.ViewProblems()
	return
}

func UpdateProblem(p model.Problem) (err error) {
	err = dao.UpdateProduct(p)
	return
}
