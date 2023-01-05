package request_html

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/millirud/go-scrapper/internal/entity"
	"github.com/rs/zerolog"
)

func New(
	logger *zerolog.Logger,
	timeout time.Duration,
	retryCount int,
) *RequestHtmlService {

	client := resty.New().
		SetTimeout(timeout).
		SetRetryCount(retryCount)

	return &RequestHtmlService{
		logger: logger,
		client: client,
	}
}

type RequestHtmlService struct {
	logger *zerolog.Logger
	client *resty.Client
}

func (r *RequestHtmlService) Do(url string) (*entity.HtmlPage, error) {

	r.logger.Info().Str("url", url).Msg("RequestHtmlService.Do.StartRequest")

	resp, err := r.client.R().
		Get(url)

	if err != nil {
		r.logger.Err(err).Msg("RequestHtmlService.Do.Error")
		return nil, err
	}

	return &entity.HtmlPage{
		Html:       resp.String(),
		StatusCode: resp.StatusCode(),
	}, nil

}
