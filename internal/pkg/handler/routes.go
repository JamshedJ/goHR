package handler

import (
	"github.com/gin-gonic/gin"
)

func (s *server) initRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(mwLogRequests)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", s.signUp)
		auth.POST("/sign-in", s.signIn)
	}

	user := router.Group("/user", s.mwUserAuth)
	{
		user.GET("/:id", s.getUserById)
		user.GET("/", s.getAllUsers)
		user.PUT("/", mwAdmin, s.updateUser)
		user.DELETE("/", mwAdmin, s.deleteUser)
	}

	employee := router.Group("/employee", s.mwUserAuth)
	{
		employee.POST("/", mwAdmin, s.createEmployee)
		employee.GET("/:id", s.getEmployeeByID)
		employee.GET("/", s.getAllEmployees)
		employee.PUT("/:id", mwAdmin, s.updateEmployee)
		employee.DELETE("/:id", mwAdmin, s.deleteEmployee)
		employee.GET("/search", s.searchEmployeeByName)
	}

	department := router.Group("/department", s.mwUserAuth)
	{
		department.POST("/", mwAdmin, s.createDepartment)
		department.GET("/", s.getAllDepartments)
		department.GET("/:id", mwAdmin, s.getDepartmentByID)
		department.PUT("/:id", mwAdmin, s.updateDepartment)
		department.DELETE("/:id", mwAdmin, s.deleteDepartment)
	}

	position := router.Group("/position", s.mwUserAuth)
	{
		position.POST("/", mwAdmin, s.createPosition)
		position.GET("/", s.getAllPositions)
		position.GET("/:id", s.getPositionByID)
		position.PUT("/:id", mwAdmin, s.updatePosition)
		position.DELETE("/:id", mwAdmin, s.deletePosition)
	}

	employeeRequest := router.Group("/employee_request", s.mwUserAuth)
	{
		employeeRequest.POST("/", mwAdmin, s.createEmployeeRequest)
		employeeRequest.GET("/", s.getAllEmployeeRequests)
		employeeRequest.GET("/:id", s.getEmployeeRequestByID)
		employeeRequest.PUT("/:id", mwAdmin, s.updateEmployeeRequest)
		employeeRequest.DELETE("/:id", mwAdmin, s.deleteEmployeeRequest)
	}

	employeeRequestTypes := router.Group("/employee_request_type", s.mwUserAuth)
	{
		employeeRequestTypes.POST("/", mwAdmin, s.createEmployeeRequestType)
		employeeRequestTypes.GET("/", s.getAllEmployeerequestTypes)
		employeeRequestTypes.GET("/:id", s.getEmployeeRequestTypeByID)
		employeeRequestTypes.PUT("/:id", mwAdmin, s.updateEmployeeRequestType)
		employeeRequestTypes.DELETE("/:id", mwAdmin, s.deleteEmployeeRequestType)
	}

	return router
}
