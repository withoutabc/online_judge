package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
)

type SubmissionServiceImpl struct {
	SubmissionService
}

func NewSubmissionApi() *SubmissionServiceImpl {
	return &SubmissionServiceImpl{
		SubmissionService: service.NewSubmissionServiceImpl(),
	}
}

type SubmissionService interface {
	AddSubmission(submission model.Submission) int
	SearchSubmission(req model.ReqSearchSubmission) ([]model.Submission, int)
}

func (s *SubmissionServiceImpl) Submit(c *gin.Context) {
	//receive
	var submission model.Submission
	if err := c.ShouldBind(&submission); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := s.SubmissionService.AddSubmission(submission)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func (s *SubmissionServiceImpl) SearchSubmission(c *gin.Context) {
	var req model.ReqSearchSubmission
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	submissions, code := s.SubmissionService.SearchSubmission(req)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	}
	util.RespNormSuccess(c, submissions)
}
