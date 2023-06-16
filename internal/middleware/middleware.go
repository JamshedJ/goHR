package middleware

import (
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MwGetID(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	c.Set("id", id)
}

//func (s *server) mwUserAuth(c *gin.Context) {
//	token := c.Request.Header.Get("token")
//	if token == "" {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Unauthorized)
//		return
//	}
//	id, err := s.app.ParseToken(token)
//	if err != nil {
//		models.ReplyError(c, err)
//		c.Abort()
//		return
//	}
//	c.Set("user_id", id)
//}
