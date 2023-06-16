package handler

import (
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getUserById(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	user, err := s.app.GetUserById(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) getAllUsers(c *gin.Context) {
	users, err := s.app.GetAllUsers(c.Request.Context())
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) updateUser(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	err := s.app.UpdateUser(c.Request.Context(), id, userID, u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *Server) deleteUser(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	err := s.app.DeleteUser(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
