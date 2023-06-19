package handler

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

	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	id, err := s.app.CreateEmployee(c.Request.Context(), u, e)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getEmployeeByID(c *gin.Context) {
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	employee, err := s.app.GetEmployeeByID(c.Request.Context(), u, id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (s *server) getEmployees(c *gin.Context) {
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	employees, err := s.app.GetEmployees(c.Request.Context(), u)
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
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}
	e.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.UpdateEmployee(c.Request.Context(), u, e)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteEmployee(c *gin.Context) {
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.DeleteEmployee(c.Request.Context(), u, id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
