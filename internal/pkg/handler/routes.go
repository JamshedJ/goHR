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

	employeeRequest := router.Group("/employee_request")
	{
		employeeRequest.POST("/", s.createEmployeeRequest)
		employeeRequest.GET("/", s.getAllEmployeeRequests)
		employeeRequest.GET("/:id", s.getEmployeeRequestByID)
		employeeRequest.PUT("/:id", s.updateEmployeeRequest)
		employeeRequest.DELETE("/:id", s.deleteEmployeeRequest)
	}

	employeeRequestTypes := router.Group("/employee_request_type", s.mwUserAuth)
	{
		employeeRequestTypes.POST("/", s.createEmployeeRequestType)
		employeeRequestTypes.GET("/", s.getAllEmployeerequestTypes)
		employeeRequestTypes.GET("/:id", s.getEmployeeRequestTypeByID)
		employeeRequestTypes.PUT("/:id", s.updateEmployeeRequestType)
		employeeRequestTypes.DELETE("/:id", s.deleteEmployeeRequestType)
	}

	return router
}
