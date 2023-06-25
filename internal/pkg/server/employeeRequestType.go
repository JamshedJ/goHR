package server

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createEmployeeRequestType(c *gin.Context) {
	var e models.EmployeeRequestType
	if err := c.BindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.service.CreateEmployeeRequestType(c.Request.Context(), e)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getEmployeeRequestTypeByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	erType, err := s.service.GetEmployeeRequestTypeByID(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, erType)
}

func (s *server) getAllEmployeerequestTypes(c *gin.Context) {
	employeeRequestTypes, err := s.service.GetAllEmployeeRequestTypes(c.Request.Context())
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, employeeRequestTypes)
}

func (s *server) updateEmployeeRequestType(c *gin.Context) {
	var t models.EmployeeRequestType
	err := c.BindJSON(&t)
	if err != nil {
		replyError(c, err)
		return
	}

	t.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.UpdateEmployeeRequestType(c.Request.Context(), t)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteEmployeeRequestType(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.DeleteEmployeeRequestType(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
