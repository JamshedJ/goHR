package handler

// import (
// 	"net/http"

// 	"github.com/JamshedJ/goHR/internal/models"
// 	"github.com/gin-gonic/gin"
// )

// func (s *server) createRequest(c *gin.Context) {
// 	var r models.Request
// 	if err := c.BindJSON(&r); err != nil {
// 		c.JSON(http.StatusBadRequest, models.BadRequest)
// 		return
// 	}

// 	id, err := s.app.createRequest(c.Request.Context(), r)
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
// }

// func (s *server) getRequestByID(c *gin.Context) {
// 	id, err := getParamInt(c, "id")
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}

// 	employee, err := s.app.GetRequestByID(c.Request.Context(), id)
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, employee)
// }

// func (s *server) getAllRequests(c *gin.Context) {
// 	requests, err := s.app.GetAllRequests(c.Request.Context())
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, requests)
// }

// func (s *server) updateRequest(c *gin.Context) {
// 	var r models.Request
// 	err := c.BindJSON(&r)
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
	
// 	r.ID, err = getParamInt(c, "id")
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}

// 	err = s.app.UpdateRequest(c.Request.Context(), r)
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, models.OK)
// }

// func (s *server) deleteRequest(c *gin.Context) {
// 	id, err := getParamInt(c, "id")
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}

// 	err = s.app.DeleteRequest(c.Request.Context(), id)
// 	if err != nil {
// 		replyError(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, models.OK)
// }
