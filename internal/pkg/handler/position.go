package handler

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

	id, err := s.app.CreatePosition(c.Request.Context(), p)
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

	position, err := s.app.GetPositionByID(c.Request.Context(), id, isAdmin)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, position)
}

func (s *server) getAllPositions(c *gin.Context) {
<<<<<<< HEAD
	positions, err := s.app.GetAllPositions(c.Request.Context())
=======
	_, isAdmin := c.Get("is_admin")

	positions, err := s.app.GetAllPositions(c.Request.Context(), isAdmin)
>>>>>>> 53f4c62490a669c7617f897c1e2e393dd2e02c36
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
		replyError(c, err)
		return
	}

	p.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.UpdatePosition(c.Request.Context(), p)
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

	err = s.app.DeletePosition(c.Request.Context(), id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
