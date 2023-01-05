package request_html

import (
	"time"

	"github.com/millirud/go-scrapper/internal/entity"
	"github.com/rs/zerolog"
)

func New(
	logger *zerolog.Logger,
	timeout time.Duration,
) *RequestHtmlService {
	return &RequestHtmlService{
		logger:  logger,
		Timeout: timeout,
	}
}

type RequestHtmlService struct {
	logger  *zerolog.Logger
	Timeout time.Duration
}

func (r *RequestHtmlService) Do(url string) (*entity.HtmlPage, error) {

}
