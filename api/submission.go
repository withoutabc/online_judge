package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
)

func Submit(c *gin.Context) {
	uid := c.Param("uid")
	//获取提交的信息
	s := model.Submission{
		Pid:      c.PostForm("pid"),
		Uid:      uid,
		Code:     c.PostForm("code"),
		Language: c.PostForm("language"),
	}
	//全部不能为空
	if s.Pid == "" || s.Uid == "" || s.Code == "" || s.Language == "" {
		util.RespParamErr(c)
		return
	}
	//根据uid得到submissions切片，同一用户的code不得重复提交
	s1 := model.Submission{
		Uid: s.Uid,
	}
	submissions, err := service.ViewResult(s1)
	if err != nil {
		fmt.Printf("view result err:%v", err)
		util.RespInternalErr(c)
		return
	}
	for _, submission := range submissions {
		if submission.Code == s.Code && submission.Language == s.Language {
			util.NormErr(c, 400, "same submission")
			return
		}
	}
	//写入提交的数据
	err = service.Submit(s)
	if err != nil {
		fmt.Printf("submit submission err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "submit success")
}

func ViewResult(c *gin.Context) {
	uid := c.Param("uid")
	//获取查询字段，都可为空
	s := model.Submission{
		Pid:      c.PostForm("pid"),
		Uid:      uid,
		Language: c.PostForm("language"),
		Status:   c.PostForm("status"),
	}
	//查询
	submissions, err := service.ViewResult(s)
	if err != nil {
		fmt.Printf("view result err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespSubmission{
		Status: 200,
		Info:   "view result success",
		Data:   submissions,
	})
}
