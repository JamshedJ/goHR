package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

type server struct {
	service *service.Service
}

func Run(ctx context.Context, service *service.Service, addr string) error {
	srv := &server{
		service: service,
	}

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      srv.initRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownErrorChan := make(chan error)

	go func() {
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan

		ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		shutdownErrorChan <- httpServer.Shutdown(ctxTimeout)
	}()

	err := httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	return nil
}

func replyError(c *gin.Context, err error) {
	switch err.(type) {
	case models.ErrorBadRequest:
		c.JSON(http.StatusBadRequest, err)
	default:
		switch err {
		case models.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, models.Unauthorized)
		// case models.ErrBadRequest:
		// 	c.JSON(http.StatusBadRequest, models.BadRequest)
		case models.ErrNoRows:
			c.JSON(http.StatusNotFound, models.NotFound)
		case models.ErrDuplicate:
			c.JSON(http.StatusNotAcceptable, models.Duplicate)
		default:
			log.Error.Println("http replyError unhandled error:", err)
			c.JSON(http.StatusInternalServerError, models.InternalErr)
		}
	}
	return
}

func getParamInt(c *gin.Context, param string) (val int, err error) {
	idStr, ok := c.Params.Get(param)
	if !ok {
		err = models.NewErrorBadRequest("no param provided")
		return
	}
	val, err = strconv.Atoi(idStr)
	if err != nil {
		err = models.NewErrorBadRequest("invalid param")
		return
	}
	return
}
