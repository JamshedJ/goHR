package server

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func mwLogRequests(c *gin.Context) {
	requestBody, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(requestBody))

	responseWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = responseWriter

	log.Debug.Printf("Request: %s", requestBody)
	c.Next()
	log.Debug.Printf("Response: %s", responseWriter.body.Bytes())
}

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

	user, err := s.service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Unauthorized)
		return
	}
	c.Set("user", user)
	c.Set("is_admin", user.IsAdmin())
}

// всё это, или...

func mwAdmin(c *gin.Context) {
	if _, ok := c.Get("is_admin"); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.InternalErr)
	}
}
