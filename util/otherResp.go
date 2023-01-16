package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_judge/model"
)

func ViewProblems(c *gin.Context, info string, problems []model.Problem) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   problems,
	})
}

func ViewSubmissions(c *gin.Context, info string, submissions []model.Submission) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   submissions,
	})
}

func ViewTestcases(c *gin.Context, info string, testcases []model.Testcase) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   testcases,
	})
}
