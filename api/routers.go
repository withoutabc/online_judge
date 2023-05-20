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
		u.POST("/password/:user_id", middleware.JWTAuthMiddleware(), uapi.ChangePassword)
	}

	p := r.Group("/problem")
	{
		papi := NewProblemApi()
		p.POST("/add", middleware.JWTAuthMiddleware(), papi.AddProblem)
		p.GET("/search", papi.SearchProblem)
		p.PUT("/update/:problem_id", middleware.JWTAuthMiddleware(), papi.UpdateProblem)
		p.DELETE("/delete/:problem_id", middleware.JWTAuthMiddleware(), papi.DeleteProblem)
	}
	s := r.Group("/submission")
	{
		sapi := NewSubmissionApi()
		s.POST("/submit", middleware.JWTAuthMiddleware(), sapi.Submit)
		s.GET("/search", middleware.JWTAuthMiddleware(), sapi.SearchSubmission)
	}
	t := r.Group("/test")
	{
		t.Use(middleware.JWTAuthMiddleware())
		sapi := NewTestApi()
		t.POST("/add", sapi.AddTestcase)
		t.GET("/search/:problem_id", sapi.SearchTestcase)
		t.PUT("/update/:testcase_id", sapi.UpdateTestcase)
		t.DELETE("/delete/:testcase_id", sapi.DeleteTestcase)
	}

	r.GET("/ranking")
	r.Run(":2333")
}
