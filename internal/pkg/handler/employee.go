package handler

import (
	"fmt"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) createEmployee(c *gin.Context) {
	var e models.Employee
	if err := c.BindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	//e.ID = c.GetInt("id")

	id, err := s.app.CreateEmployee(c.Request.Context(), e)
	if err != nil {
		fmt.Println("first case")
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *Server) getEmployeeByID(c *gin.Context) {
	id := c.GetInt("id")
	employee, err := s.app.GetEmployeeByID(c.Request.Context(), id)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}