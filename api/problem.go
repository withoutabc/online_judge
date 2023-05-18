package api

import (
	"github.com/gin-gonic/gin"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

type ProblemServiceImpl struct {
	ProblemService
}

func NewProblemApi() *ProblemServiceImpl {
	return &ProblemServiceImpl{
		ProblemService: service.NewProblemServiceImpl(),
	}
}

type ProblemService interface {
	AddProblem(problem model.Problem) int
	SearchProblem(request model.ReqSearchProblem) (problems []model.Problem, n int)
	UpdateProblem(problemId int64, problem model.Problem) int
	DeleteProblem(problemId int64) int
}

func (p *ProblemServiceImpl) AddProblem(c *gin.Context) {
	//获取题目信息
	var problem model.Problem
	if err := c.ShouldBind(&problem); err != nil {
		util.Find()
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := p.ProblemService.AddProblem(problem)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func (p *ProblemServiceImpl) SearchProblem(c *gin.Context) {
	var request model.ReqSearchProblem
	if err := c.ShouldBind(&request); err != nil {
		util.Find()
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	exist, _ := util.InArray(request.Level, []string{"极易", "容易", "中等", "困难", "极难"})
	if !exist {
		util.Find()
		util.NormErr(c, util.WrongLevelCode)
		return
	}
	if util.CheckTime(request.From, request.To) != nil {
		util.Find()
		util.NormErr(c, util.WrongTimeCode)
		return
	}
	problems, code := p.ProblemService.SearchProblem(request)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	}
	util.RespNormSuccess(c, problems)
}

func (p *ProblemServiceImpl) UpdateProblem(c *gin.Context) {
	//receive
	problemId := c.Param("problem_id")
	IntProblemId, err := strconv.ParseInt(problemId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	var problem model.Problem
	if err = c.ShouldBind(&problem); err != nil {
		util.Find()
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := p.ProblemService.UpdateProblem(IntProblemId, problem)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func (p *ProblemServiceImpl) DeleteProblem(c *gin.Context) {
	//receive
	problemId := c.Param("problem_id")
	IntProblemId, err := strconv.ParseInt(problemId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	code := p.ProblemService.DeleteProblem(IntProblemId)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

//func SearchProblem(c *gin.Context) {
//	pid := c.Query("pid")
//	//查看题目
//	problems, err := service.SearchProblems(pid)
//	if err != nil {
//		fmt.Printf("search problem err:%v", err)
//		util.RespInternalErr(c)
//		return
//	}
//	c.JSON(http.StatusOK, model.RespProblem{
//		Status: 200,
//		Info:   "search problem success",
//		Data:   problems,
//	})
//}
//
//func UpdateProblem(c *gin.Context) {
//	//获取信息
//	uid := c.Param("uid")
//	Pid := c.PostForm("pid")
//	if uid == "" || Pid == "" {
//		util.RespParamErr(c)
//		return
//	}
//	pid, err := strconv.Atoi(Pid)
//	if err != nil {
//		fmt.Println(err)
//		util.NormErr(c, 400, "invalid pid")
//		return
//	}
//	timeLimit := c.PostForm("time_limit")
//	memoryLimit := c.PostForm("memory_limit")
//	p := model.Problem{
//		Pid:               pid,
//		Title:             c.PostForm("title"),
//		Description:       c.PostForm("description"),
//		DescriptionInput:  c.PostForm("description_input"),
//		DescriptionOutput: c.PostForm("description_output"),
//		SampleInput:       c.PostForm("sample_input"),
//		SampleOutput:      c.PostForm("sample_output"),
//		Uid:               uid,
//	}
//	//都不填为更新失败
//	if p.Title == "" && p.Description == "" && p.DescriptionInput == "" && p.DescriptionOutput == "" && p.SampleInput == "" && p.SampleOutput == "" && timeLimit == "" && memoryLimit == "" {
//		util.NormErr(c, 400, "fail to update")
//		return
//	}
//	t, err := strconv.ParseFloat(timeLimit, 64)
//	if timeLimit != "" && err != nil {
//		util.NormErr(c, 400, "invalid time limit")
//		return
//	}
//	var m float64
//	m, err = strconv.ParseFloat(memoryLimit, 64)
//	if memoryLimit != "" && err != nil {
//		util.NormErr(c, 400, "invalid memory limit")
//		return
//	}
//	p.TimeLimit = t
//	p.MemoryLimit = m
//	problems, err := service.SearchProblems(Pid)
//	if err != nil {
//		fmt.Printf("search problems err:%v", err)
//		util.RespInternalErr(c)
//		return
//	}
//	//题号不存在
//	if problems == nil {
//		fmt.Printf("problems:nil\n")
//		util.NormErr(c, 400, "pid not exist")
//		return
//	}
//	//是否出现了相同的title
//	var problems1 []model.Problem
//	problems1, err = service.SearchProblems("")
//	if err != nil {
//		fmt.Printf("view problems err:%v", err)
//		util.RespInternalErr(c)
//		return
//	}
//	for _, problem := range problems1 {
//		if problem.Title == p.Title {
//			util.NormErr(c, 400, "same title")
//			return
//		}
//	}
//	//修改后信息都重复
//	if service.CheckStruct(problems, p) {
//		util.NormErr(c, 400, "repeated problem")
//		return
//	}
//	//修改题目
//	err = service.UpdateProblem(p)
//	if err != nil {
//		fmt.Printf("update problem err:%v", err)
//		util.RespInternalErr(c)
//		return
//	}
//	util.RespOK(c, "update problem success")
//}
