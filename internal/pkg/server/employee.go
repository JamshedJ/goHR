package server

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createEmployee(c *gin.Context) {
	var e models.Employee
	if err := c.BindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.service.CreateEmployee(c.Request.Context(), e)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getEmployeeByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	_, isAdmin := c.Get("is_admin")

	employee, err := s.service.GetEmployeeByID(c.Request.Context(), id, isAdmin)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (s *server) getAllEmployees(c *gin.Context) {
	_, isAdmin := c.Get("is_admin")

	employees, err := s.service.GetAllEmployees(c.Request.Context(), isAdmin)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (s *server) updateEmployee(c *gin.Context) {
	var e models.Employee
	err := c.BindJSON(&e)
	if err != nil {
		replyError(c, err)
		return
	}
	e.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.UpdateEmployee(c.Request.Context(), e)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteEmployee(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.DeleteEmployee(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) searchEmployeeByName(c *gin.Context) {
	query := c.Query("query")

	employee, err := s.service.SearchEmployeeByName(c.Request.Context(), query)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}
