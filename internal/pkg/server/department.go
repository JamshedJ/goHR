package server

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createDepartment(c *gin.Context) {
	var d models.Department
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.service.CreateDepartment(c.Request.Context(), d)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getDepartmentByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	department, err := s.service.GetDepartmentByID(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, department)
}

func (s *server) getAllDepartments(c *gin.Context) {
	departments, err := s.service.GetAllDepartments(c.Request.Context())
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, departments)
}

func (s *server) updateDepartment(c *gin.Context) {
	var d models.Department
	err := c.BindJSON(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	d.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.UpdateDepartment(c.Request.Context(), d)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteDepartment(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.DeleteDepartment(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
