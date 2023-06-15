package handler

import (
	"context"
	"errors"
	"github.com/JamshedJ/goHR"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	app *goHR.App
}

func Run(ctx context.Context, app *goHR.App, addr string) error {
	srv := &Server{
		app: app,
	}

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      srv.InitRoutes(),
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
