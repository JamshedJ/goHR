package handler

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

	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	id, err := s.app.CreateEmployeeRequest(c.Request.Context(), u, r)
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

	employee, err := s.app.GetEmployeeRequestByID(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (s *server) getAllEmployeeRequests(c *gin.Context) {
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	requests, err := s.app.GetAllEmployeeRequests(c.Request.Context(), u)
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

	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.UpdateEmployeeRequest(c.Request.Context(), u, r)
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

	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.DeleteEmployeeRequest(c.Request.Context(), u, id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
