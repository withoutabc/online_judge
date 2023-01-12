package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

func AddProblem(c *gin.Context, p model.Problem) {
	err := dao.InsertProblem(p)
	if err != nil {
		fmt.Printf("add product err:%v", err)
		util.RespInternalErr(c)
		return
	}
}

func ViewProblems(c *gin.Context) (problems []model.Problem) {
	var err error
	problems, err = dao.ViewProblems()
	if err != nil {
		fmt.Printf("view problems err:%v", err)
		util.RespInternalErr(c)
		return nil
	}
	return
}

func UpdateProblem(c *gin.Context, p model.Problem) {
	err := dao.UpdateProduct(p)
	if err != nil {
		fmt.Printf("update problem err:%v", err)
		util.RespInternalErr(c)
		return
	}
}
