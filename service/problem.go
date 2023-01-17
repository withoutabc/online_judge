package service

import (
	"online_judge/dao"
	"online_judge/model"
)

func AddProblem(p model.Problem) (err error) {
	err = dao.InsertProblem(p)
	return
}

func SearchProblems(pid string) (problems []model.Problem, err error) {
	problems, err = dao.SearchProblems(pid)
	return
}

func UpdateProblem(p model.Problem) (err error) {
	err = dao.UpdateProblem(p)
	return
}

func CheckStruct(s []model.Problem, existing model.Problem) bool {
	for _, v := range s {
		match := true
		if existing.Title != "" && v.Title != existing.Title {
			match = false
		}
		if existing.Description != "" && v.Description != existing.Description {
			match = false
		}
		if existing.DescriptionInput != "" && v.DescriptionInput != existing.DescriptionInput {
			match = false
		}
		if existing.DescriptionOutput != "" && v.DescriptionOutput != existing.DescriptionOutput {
			match = false
		}
		if existing.SampleInput != "" && v.SampleInput != existing.SampleInput {
			match = false
		}
		if existing.SampleOutput != "" && v.SampleOutput != existing.SampleOutput {
			match = false
		}
		if existing.TimeLimit != 0.0 && v.TimeLimit != existing.TimeLimit {
			match = false
		}
		if existing.MemoryLimit != 0.0 && v.MemoryLimit != existing.MemoryLimit {
			match = false
		}
		if match {
			return true
		}
	}
	return false
}
