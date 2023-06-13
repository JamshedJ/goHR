package handler

import (
	"github.com/JamshedJ/goHR/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	employee := router.Group("/employee")
	{
		employee.POST("/")
		employee.GET("/")
		employee.GET("/:id")
		employee.PUT("/:id")
		employee.DELETE("/:id")
	}

	department := router.Group("/department")
	{
		department.POST("/")
		department.GET("/")
		department.GET("/:id")
		department.PUT("/:id")
		department.DELETE("/:id")
	}

	vacancy := router.Group("/vacancy")
	{
		vacancy.POST("/")
		vacancy.GET("/")
		vacancy.GET("/:id")
		vacancy.PUT("/:id")
		vacancy.DELETE("/:id")
	}

	users := router.Group("/users")
	{
		users.GET("/:id")
		users.GET("/")
		users.PUT("/")
		users.DELETE("/")
	}

	return router
}
