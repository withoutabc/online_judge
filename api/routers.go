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
		u.GET("/refresh", Refresh)
		u.POST("/logout", middleware.Auth(), Logout)
		u.POST("/password", middleware.Auth(), ChangePassword)
	}
	p := r.Group("/problem")
	{
		p.POST("/add", middleware.Auth(), AddProblem)
		p.GET("/view", ViewProblem)
		p.PUT("/update", middleware.Auth(), UpdateProblem)
	}
	s := r.Group("/submission")
	{
		s.POST("/submit", middleware.Auth(), Submit)
		s.GET("/view", ViewResult)
	}
	r.Run()
}
