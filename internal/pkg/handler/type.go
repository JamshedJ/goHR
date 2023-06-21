package handler

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createType(c *gin.Context) {
	var t models.Type
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.app.CreateType(c.Request.Context(), t)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getTypeByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	types, err := s.app.GetTypeByID(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, types)
}

func (s *server) getAllTypes(c *gin.Context) {
	types, err := s.app.GetAllTypes(c.Request.Context())
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, types)
}

func (s *server) updateType(c *gin.Context) {
	var t models.Type
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

	err = s.app.UpdateType(c.Request.Context(), t)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteType(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.DeleteType(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
