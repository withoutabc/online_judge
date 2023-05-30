package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

type TestServiceImpl struct {
	TestService
}

func NewTestApi() *TestServiceImpl {
	return &TestServiceImpl{
		TestService: service.NewTestServiceImpl(),
	}
}

type TestService interface {
	AddTestcase(testcase model.Testcase) (model.Testcase, int)
	SearchTestcase(problemId int64) ([]model.Testcase, int)
	UpdateTestcase(testcaseId int64, testcase model.Testcase) int
	DeleteTestcase(testcaseId int64) int
}

func (t *TestServiceImpl) AddTestcase(c *gin.Context) {
	var testcase model.Testcase
	if err := c.ShouldBind(&testcase); err != nil {
		log.Println(err)
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	testcase, code := t.TestService.AddTestcase(testcase)
	switch code {
	case util.InternalServeErrCode:
		util.NormErr(c, util.InternalServeErrCode)
		return
	}
	util.RespNormSuccess(c, testcase)
}

func (t *TestServiceImpl) SearchTestcase(c *gin.Context) {
	problemId := c.Param("problem_id")
	IntProblemId, err := strconv.ParseInt(problemId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	testcases, code := t.TestService.SearchTestcase(IntProblemId)
	switch code {
	case util.InternalServeErrCode:
		util.NormErr(c, util.InternalServeErrCode)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	}
	util.RespNormSuccess(c, testcases)
}

func (t *TestServiceImpl) UpdateTestcase(c *gin.Context) {
	testcaseId := c.Param("testcase_id")
	IntTestcaseId, err := strconv.ParseInt(testcaseId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	var testcase model.Testcase
	if err = c.ShouldBind(&testcase); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := t.TestService.UpdateTestcase(IntTestcaseId, testcase)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.UpdateFailErrCode:
		util.NormErr(c, util.UpdateFailErrCode)
		return
	}
	util.RespOK(c)
}

func (t *TestServiceImpl) DeleteTestcase(c *gin.Context) {
	testcaseId := c.Param("testcase_id")
	IntTestcaseId, err := strconv.ParseInt(testcaseId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	code := t.TestService.DeleteTestcase(IntTestcaseId)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.UpdateFailErrCode:
		util.NormErr(c, util.UpdateFailErrCode)
		return
	}
	util.RespOK(c)
}
