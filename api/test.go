package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
)

func AddTestcase(c *gin.Context) {
	uid := c.Param("uid")
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
	//输入或输入不得重复
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
	util.RespOK(c)
}

func ViewTestcases(c *gin.Context) {
	uid := c.Param("uid")
	//获取pid
	pid := c.Query("pid")
	testcases, err := service.SearchTestcase(uid, pid)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.ViewTestcases(c, "view testcases successfully", testcases)
}

func UpdateTestcase(c *gin.Context) {
	uid := c.Param("uid")
	//获取修改信息
	pid := c.PostForm("pid")
	input := c.PostForm("input")
	output := c.PostForm("output")
	t := model.Testcase{
		Uid:    uid,
		Pid:    pid,
		Input:  input,
		Output: output,
	}
	//pid不能为空
	if pid == "" {
		util.RespParamErr(c)
		return
	}
	//input output 不能都为空
	if input == "" && output == "" {
		util.NormErr(c, 400, "fail to update")
		return
	}
	//input和output不能同时重复
	testcases, err := service.SearchTestcase(uid, pid)
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
	//执行修改
	err = service.UpdateTestcase(t)
	if err != nil {
		fmt.Printf("update testcase err:%v", err)
		util.RespInternalErr(c)
		return
	}
}

func DeleteTestcase(c *gin.Context) {
	uid := c.Param("uid")
	//tid
	tid := c.Query("tid")
	//删除
	err := service.DeleteTestcase(uid, tid)
	if err != nil {
		fmt.Printf("delete testcase err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
