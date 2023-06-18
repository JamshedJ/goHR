package handler

import (
	"github.com/JamshedJ/goHR/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", s.signUp)
		auth.POST("/sign-in", s.signIn)
	}

	employee := router.Group("/employee")
	{
		employee.POST("/", s.createEmployee)
		//employee.GET("/" )
		//employee.GET("/:id")
		//employee.PUT("/:id")
		//employee.DELETE("/:id")
	}

	//department := router.Group("/department")
	//{
	//	department.POST("/")
	//	department.GET("/")
	//	department.GET("/:id")
	//	department.PUT("/:id")
	//	department.DELETE("/:id")
	//}
	//
	//position := router.Group("/position")
	//{
	//	position.POST("/")
	//	position.GET("/")
	//	position.GET("/:id")
	//	position.PUT("/:id")
	//	position.DELETE("/:id")
	//}
	//
	//vacationRequest := router.Group("/vacationRequest")
	//{
	//	vacationRequest.POST("/")
	//	vacationRequest.GET("/")
	//	vacationRequest.GET("/:id")
	//	vacationRequest.PUT("/:id")
	//	vacationRequest.DELETE("/:id")
	//}
	//
	//sickLeaveRequest := router.Group("/sickLeaveRequest")
	//{
	//	sickLeaveRequest.POST("/")
	//	sickLeaveRequest.GET("/")
	//	sickLeaveRequest.GET("/:id")
	//	sickLeaveRequest.PUT("/:id")
	//	sickLeaveRequest.DELETE("/:id")
	//}

	user := router.Group("/user", middleware.MwGetID)
	{
		user.GET("/:id", s.getUserById)
		user.GET("/", s.getAllUsers)
		user.PUT("/", s.updateUser)
		user.DELETE("/", s.deleteUser)
	}

	return router
}
