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
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := p.ProblemService.AddProblem(problem)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.RepeatedTitleErrCode:
		util.NormErr(c, util.RepeatedTitleErrCode)
		return
	}
	util.RespOK(c)
}

func (p *ProblemServiceImpl) SearchProblem(c *gin.Context) {
	var request model.ReqSearchProblem
	if err := c.ShouldBindJSON(&request); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	if util.CheckTime(request.From, request.To) != nil {
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
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := p.ProblemService.UpdateProblem(IntProblemId, problem)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	case util.UpdateFailErrCode:
		util.NormErr(c, util.UpdateFailErrCode)
		return
	case util.RepeatedTitleErrCode:
		util.NormErr(c, util.RepeatedTitleErrCode)
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
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	case util.UpdateFailErrCode:
		util.NormErr(c, util.UpdateFailErrCode)
		return
	}
	util.RespOK(c)
}
