package handler

import (
	"net/http"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) getUserById(c *gin.Context) {
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

	user, err := s.app.GetUserById(c.Request.Context(), u, id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *server) getAllUsers(c *gin.Context) {
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}
	users, err := s.app.GetAllUsers(c.Request.Context(), u)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *server) updateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	user.ID, err = getParamInt(c, "id")
	if err != nil {
		replyError(c, err)
		return
	}
	u, err := getUserFromContext(c)
	if err != nil {
		replyError(c, err)
		return
	}

	err = s.app.UpdateUser(c.Request.Context(), u, user)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (s *server) deleteUser(c *gin.Context) {
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

	err = s.app.DeleteUser(c.Request.Context(), u, id)
	if err != nil {
		replyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
