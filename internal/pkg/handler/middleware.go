package handler

import (
	"net/http"
	"strings"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *server) mwUserAuth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Unauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "token is empty"})
		return
	}

	user, err := s.app.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Unauthorized)
		return
	}
	c.Set("user", user)
}
