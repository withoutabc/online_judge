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

func AddTestcase(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		util.RespParamErr(c)
		return
	}
	t := model.Testcase{
		Pid:    c.PostForm("pid"),
		Uid:    uid,
		Input:  c.PostForm("input"),
		Output: c.PostForm("output"),
	}
	if t.Pid == "" || t.Input == "" || t.Output == "" {
		util.RespParamErr(c)
		return
	}
	//输入或输出不得重复
	testcases, err := service.SearchTestcase(uid, t.Pid)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	for _, testcase := range testcases {
		if testcase.Input == t.Input && testcase.Output == t.Output {
			util.NormErr(c, 400, "same or wrong testcase")
			return
		}
	}
	//添加
	err = service.AddTestcase(t)
	if err != nil {
		fmt.Printf("add testcase err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "add testcase success")
}

func ViewTestcases(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		util.RespParamErr(c)
		return
	}
	//获取pid
	pid := c.Query("pid")
	testcases, err := service.SearchTestcase(uid, pid)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespTestcase{
		Status: 200,
		Info:   "view testcase success",
		Data:   testcases,
	})
}

func UpdateTestcase(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		util.RespParamErr(c)
		return
	}
	//获取修改信息
	pid := c.PostForm("pid")
	Tid := c.PostForm("tid")
	input := c.PostForm("input")
	output := c.PostForm("output")
	//不能为空
	if pid == "" || Tid == "" {
		util.RespParamErr(c)
		return
	}
	tid, err := strconv.Atoi(Tid)
	if err != nil {
		util.NormErr(c, 400, "invalid tid")
		return
	}
	t := model.Testcase{
		Uid:    uid,
		Pid:    pid,
		Tid:    tid,
		Input:  input,
		Output: output,
	}

	//input output 不能都为空
	if input == "" && output == "" {
		util.NormErr(c, 400, "fail to update")
		return
	}
	//input和output不能同时重复
	testcases, err := service.SearchTestcase(uid, pid)
	if err != nil {
		fmt.Printf("search testcase error:%v", err)
		util.RespInternalErr(c)
		return
	}
	for _, testcase := range testcases {
		if testcase.Input == t.Input && testcase.Output == t.Output {
			util.NormErr(c, 400, "same or wrong testcase")
			return
		}
	}
	//执行修改
	err = service.UpdateTestcase(t)
	if err != nil {
		fmt.Printf("update testcase err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update testcase success")
}

func DeleteTestcase(c *gin.Context) {
	uid := c.Param("uid")
	tid := c.Query("tid")
	if uid == "" || tid == "" {
		util.RespParamErr(c)
		return
	}
	//删除
	err := service.DeleteTestcase(uid, tid)
	if err != nil {
		fmt.Printf("delete testcase err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "delete testcase success")
}
