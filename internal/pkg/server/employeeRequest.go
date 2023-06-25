package server

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createEmployeeRequest(c *gin.Context) {
	var r models.EmployeeRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.service.CreateEmployeeRequest(c.Request.Context(), r)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getEmployeeRequestByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	employee, err := s.service.GetEmployeeRequestByID(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (s *server) getAllEmployeeRequests(c *gin.Context) {
	requests, err := s.service.GetAllEmployeeRequests(c.Request.Context())
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, requests)
}

func (s *server) updateEmployeeRequest(c *gin.Context) {
	var r models.EmployeeRequest
	err := c.BindJSON(&r)
	if err != nil {
		replyError(c, err)
		return
	}

	r.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.UpdateEmployeeRequest(c.Request.Context(), r)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteEmployeeRequest(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.DeleteEmployeeRequest(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
