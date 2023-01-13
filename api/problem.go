package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

func AddProblem(c *gin.Context) {
	//获取题目信息
	timeLimit := c.PostForm("time_limit")
	memoryLimit := c.PostForm("memory_limit")
	p := model.Problem{
		Title:             c.PostForm("title"),
		Description:       c.PostForm("description"),
		DescriptionInput:  c.PostForm("description_input"),
		DescriptionOutput: c.PostForm("description_output"),
		SampleInput:       c.PostForm("sample_input"),
		SampleOutput:      c.PostForm("sample_output"),
		Uid:               c.PostForm("uid"),
	}
	//所有项必填
	if p.Title == "" || p.Description == "" || p.DescriptionInput == "" || p.DescriptionOutput == "" || p.SampleInput == "" || p.SampleOutput == "" || timeLimit == "" || memoryLimit == "" {
		util.RespParamErr(c)
		return
	}
	t, err := strconv.ParseFloat(timeLimit, 64)
	if err != nil {
		util.NormErr(c, 400, "invalid time limit")
		return
	}
	var m float64
	m, err = strconv.ParseFloat(memoryLimit, 64)
	if err != nil {
		util.NormErr(c, 400, "invalid memory limit")
		return
	}
	p.TimeLimit = t
	p.MemoryLimit = m
	//是否出现了相同的title
	var problems []model.Problem
	problems, err = service.ViewProblems()
	if err != nil {
		fmt.Printf("view problems err:%v", err)
		util.RespInternalErr(c)
		return
	}
	for _, problem := range problems {
		if problem.Title == p.Title {
			util.NormErr(c, 400, "same title")
			return
		}
	}
	//插入题目
	err = service.AddProblem(p)
	if err != nil {
		fmt.Printf("add problem err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func ViewProblem(c *gin.Context) {
	//查看题目
	problems, err := service.ViewProblems()
	if err != nil {
		fmt.Printf("view problems err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.ViewProblems(c, "view problems successfully", problems)
}

func UpdateProblem(c *gin.Context) {
	//获取题目信息
	Pid := c.PostForm("pid")
	pid, err := strconv.Atoi(Pid)
	if err != nil {
		fmt.Println(err)
		util.NormErr(c, 400, "invalid pid")
		return
	}
	timeLimit := c.PostForm("time_limit")
	memoryLimit := c.PostForm("memory_limit")
	p := model.Problem{
		Pid:               pid,
		Title:             c.PostForm("title"),
		Description:       c.PostForm("description"),
		DescriptionInput:  c.PostForm("description_input"),
		DescriptionOutput: c.PostForm("description_output"),
		SampleInput:       c.PostForm("sample_input"),
		SampleOutput:      c.PostForm("sample_output"),
		Uid:               c.PostForm("uid"),
	}
	//都不填为更新失败
	if p.Title == "" && p.Description == "" && p.DescriptionInput == "" && p.DescriptionOutput == "" && p.SampleInput == "" && p.SampleOutput == "" && timeLimit == "" && memoryLimit == "" {
		util.NormErr(c, 400, "fail to update")
		return
	}
	t, err := strconv.ParseFloat(timeLimit, 64)
	if timeLimit != "" && err != nil {
		util.NormErr(c, 400, "invalid time limit")
		return
	}
	var m float64
	m, err = strconv.ParseFloat(memoryLimit, 64)
	if memoryLimit != "" && err != nil {
		util.NormErr(c, 400, "invalid memory limit")
		return
	}
	p.TimeLimit = t
	p.MemoryLimit = m
	//信息全部相同为更新失败
	var problems []model.Problem
	problems, err = service.ViewProblems()
	if err != nil {
		fmt.Printf("view problems err:%v", err)
		util.RespInternalErr(c)
		return
	}
	for _, problem := range problems {
		if problem == p {
			util.NormErr(c, 400, "fail to update")
			return
		}
	}
	//修改题目
	err = service.UpdateProblem(p)
	if err != nil {
		fmt.Printf("update problem err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
