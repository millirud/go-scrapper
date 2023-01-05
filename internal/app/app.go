// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/millirud/go-scrapper/config"
	"github.com/millirud/go-scrapper/pkg/httpserver"
	"github.com/millirud/go-scrapper/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	var err error

	l := logger.New(cfg.Log.Level, os.Stdout)

	// HTTP Server
	handler := gin.New()

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msgf("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error().Msgf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Msgf("app - Run - httpServer.Shutdown: %w", err)
	}
}
