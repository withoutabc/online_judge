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
	//s := r.Group("/submission")
	//{
	//	s.POST("/submit/:user_id", Submit)
	//	s.GET("/view/:uid", ViewResult)
	//}
	//t := r.Group("/test")
	//{
	//	t.Use(middleware.JWTAuthMiddleware())
	//	t.POST("/add/:uid", AddTestcase)
	//	t.GET("/view/:uid", ViewTestcases)
	//	t.PUT("/update/:uid", UpdateTestcase)
	//	t.DELETE("/delete/:uid", DeleteTestcase)
	//}
	r.Run(":10")
}
