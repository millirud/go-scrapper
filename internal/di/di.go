package di

import (
	"os"
	"time"

	"github.com/millirud/go-scrapper/config"
	"github.com/millirud/go-scrapper/internal/service/request_html"
	"github.com/millirud/go-scrapper/pkg/logger"
	"github.com/rs/zerolog"
)

type DI struct {
	Logger             *zerolog.Logger
	RequestHtmlService *request_html.RequestHtmlService
}

func New(cfg *config.Config) *DI {
	logger := logger.New(cfg.Log.Level, os.Stdout)

	return &DI{
		Logger: logger,
		RequestHtmlService: request_html.New(
			logger,
			time.Duration(cfg.Request.Timeout)*time.Millisecond,
			cfg.Request.RetryCount,
		),
	}
}
