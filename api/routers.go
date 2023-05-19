package api

import (
	"github.com/gin-gonic/gin"
	"online_judge/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	u := r.Group("/user")
	{
		uapi := NewUserApi()
		u.POST("/register", uapi.Register)
		u.POST("/login", uapi.Login)
		//u.POST("/refresh", middleware.JWTAuthMiddleware(), Refresh)
		u.POST("/password/:user_id", uapi.ChangePassword)
	}

	p := r.Group("/problem")
	{
		papi := NewProblemApi()
		p.POST("/add", papi.AddProblem)
		p.GET("/search", papi.SearchProblem)
		p.PUT("/update/:problem_id", papi.UpdateProblem)
		p.DELETE("/delete/:problem_id", papi.DeleteProblem)
	}
	s := r.Group("/submission")
	{
		sapi := NewSubmissionApi()
		s.POST("/submit", sapi.Submit)
		s.GET("/search", sapi.SearchSubmission)
	}
	t := r.Group("/test")
	{
		sapi := NewTestApi()
		t.POST("/add", sapi.AddTestcase)
		t.GET("/search/:problem_id", sapi.SearchTestcase)
		t.PUT("/update/:testcase_id", sapi.UpdateTestcase)
		t.DELETE("/delete/:testcase_id", sapi.DeleteTestcase)
	}
	r.Run(":2333")
}
