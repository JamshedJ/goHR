package server

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) createPosition(c *gin.Context) {
	var p models.Position
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	id, err := s.service.CreatePosition(c.Request.Context(), p)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (s *server) getPositionByID(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}
	_, isAdmin := c.Get("is_admin")

	position, err := s.service.GetPositionByID(c.Request.Context(), id, isAdmin)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, position)
}

func (s *server) getAllPositions(c *gin.Context) {
	_, isAdmin := c.Get("is_admin")

	positions, err := s.service.GetAllPositions(c.Request.Context(), isAdmin)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, positions)
}

func (s *server) updatePosition(c *gin.Context) {
	var p models.Position
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	p.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.UpdatePosition(c.Request.Context(), p)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deletePosition(c *gin.Context) {
	id, err := getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.service.DeletePosition(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
