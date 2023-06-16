package handler

import (
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) signUp(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	err := s.app.AddUser(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *Server) signIn(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	token, err := s.app.GenerateToken(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
