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
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.GET("/refresh", middleware.JWTAuthMiddleware(), Refresh)
		u.POST("/password/:uid", middleware.JWTAuthMiddleware(), ChangePassword)
	}
	p := r.Group("/problem")
	{
		p.POST("/add/:uid", middleware.JWTAuthMiddleware(), AddProblem)
		p.GET("/search", SearchProblem)
		p.PUT("/update/:uid", middleware.JWTAuthMiddleware(), UpdateProblem)
	}
	s := r.Group("/submission")
	{
		s.Use(middleware.JWTAuthMiddleware())
		s.POST("/submit/:uid", Submit)
		s.GET("/view/:uid", ViewResult)
	}
	t := r.Group("/test")
	{
		t.Use(middleware.JWTAuthMiddleware())
		t.POST("/add/:uid", AddTestcase)
		t.GET("/view/:uid", ViewTestcases)
		t.PUT("/update/:uid", UpdateTestcase)
		t.DELETE("/delete/:uid", DeleteTestcase)
	}
	r.Run()
}
