package handler

import (
	"github.com/gin-gonic/gin"
)

func (s *server) initRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", s.signUp)
		auth.POST("/sign-in", s.signIn)
	}

	user := router.Group("/user", s.mwUserAuth)
	{
		user.GET("/:id", s.getUserById)
		user.GET("/", s.getAllUsers)
		user.PUT("/", s.updateUser)
		user.DELETE("/", s.deleteUser)
	}

	employee := router.Group("/employee", s.mwUserAuth)
	{
		employee.POST("/", s.createEmployee)
		employee.GET("/:id", s.getEmployeeByID)
		employee.GET("/", s.getEmployees)
		employee.PUT("/:id", s.updateEmployee)
		employee.DELETE("/:id", s.deleteEmployee)
	}

	department := router.Group("/department")
	{
		department.POST("/", s.createDepartment)
		department.GET("/", s.getAllDepartments)
		department.GET("/:id", s.getDepartmentByID)
		department.PUT("/:id", s.updateDepartment)
		department.DELETE("/:id", s.deleteDepartment)
	}
	
	position := router.Group("/position")
	{
		position.POST("/", s.createPosition)
		position.GET("/", s.getAllPositions)
		position.GET("/:id", s.getPositionByID)
		position.PUT("/:id", s.updatePosition)
		position.DELETE("/:id", s.deletePosition)
	}
	//
	// vacationRequest := router.Group("/vacationRequest")
	// {
	//	vacationRequest.POST("/")
	//	vacationRequest.GET("/")
	//	vacationRequest.GET("/:id")
	//	vacationRequest.PUT("/:id")
	//	vacationRequest.DELETE("/:id")
	// }
	//
	// sickLeaveRequest := router.Group("/sickLeaveRequest")
	// {
	//	sickLeaveRequest.POST("/")
	//	sickLeaveRequest.GET("/")
	//	sickLeaveRequest.GET("/:id")
	//	sickLeaveRequest.PUT("/:id")
	//	sickLeaveRequest.DELETE("/:id")
	// }

	return router
}
