package api

import (
	"github.com/gin-gonic/gin"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

type InfoServiceImpl struct {
	InfoService
}

func NewInfoApi() *InfoServiceImpl {
	return &InfoServiceImpl{
		InfoService: service.NewInfoServiceImpl(),
	}
}

type InfoService interface {
	GetInfo(uid int64) (model.Info, int)
	UpdateInfo(info model.Info) int
}

func (i *InfoServiceImpl) GetInfo(c *gin.Context) {
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	info, code := i.InfoService.GetInfo(int64(IntUserId))
	switch code {
	case util.InternalServeErrCode:
		util.NormErr(c, util.InternalServeErrCode)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	}
	util.RespNormSuccess(c, info)
}

func (i *InfoServiceImpl) UpdateInfo(c *gin.Context) {
	var info model.Info
	if err := c.ShouldBind(&info); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := i.InfoService.UpdateInfo(info)
	switch code {
	case util.InternalServeErrCode:
		util.NormErr(c, util.InternalServeErrCode)
		return
	case util.UpdateFailErrCode:
		util.NormErr(c, util.UpdateFailErrCode)
		return
	}
	util.RespOK(c)
}
