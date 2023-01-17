package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

func AddProblem(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		util.RespParamErr(c)
		return
	}
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
		Uid:               uid,
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
	problems, err = service.SearchProblems("")
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
	util.RespOK(c, "add problem success")
}

func SearchProblem(c *gin.Context) {
	pid := c.Query("pid")
	//查看题目
	problems, err := service.SearchProblems(pid)
	if err != nil {
		fmt.Printf("search problem err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespProblem{
		Status: 200,
		Info:   "search problem success",
		Data:   problems,
	})
}

func UpdateProblem(c *gin.Context) {
	//获取信息
	uid := c.Param("uid")
	Pid := c.PostForm("pid")
	if uid == "" || Pid == "" {
		util.RespParamErr(c)
		return
	}
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
		Uid:               uid,
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
	problems, err := service.SearchProblems(Pid)
	if err != nil {
		fmt.Printf("search problems err:%v", err)
		util.RespInternalErr(c)
		return
	}
	if problems == nil {
		fmt.Printf("problems:nil\n")
		util.NormErr(c, 400, "pid not exist")
		return
	}
	if service.CheckStruct(problems, p) {
		util.NormErr(c, 400, "repeated problem")
		return
	}
	//修改题目
	err = service.UpdateProblem(p)
	if err != nil {
		fmt.Printf("update problem err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update problem success")
}
