package request_html

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/millirud/go-scrapper/internal/entity"
	"github.com/millirud/go-scrapper/internal/service"
	"github.com/rs/zerolog"
)

func New(
	logger *zerolog.Logger,
	timeout time.Duration,
	retryCount int,
	cookieStorage service.CookieStorage,
) *RequestHtmlService {

	client := resty.New().
		SetTimeout(timeout).
		SetRetryCount(retryCount).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return &RequestHtmlService{
		logger:        logger,
		client:        client,
		cookieStorage: cookieStorage,
	}
}

type RequestHtmlService struct {
	logger        *zerolog.Logger
	client        *resty.Client
	cookieStorage service.CookieStorage
}

func (r *RequestHtmlService) Do(url string) (*entity.HtmlPage, error) {

	r.logger.Info().Str("url", url).Msg("RequestHtmlService.Do.StartRequest")

	request := r.client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0").
		SetHeader("Accept-Language", "ru-RU").
		SetHeader("Accept-Encoding", "gzip, deflate").
		SetHeader("Accept", "text/html")

	oldCookies, err := r.cookieStorage.GetCookie(url)

	r.logger.Info().Str("url", url).Str("cookie", oldCookies).Msg("RequestHtmlService.GetCookie")

	if oldCookies != "" {
		request.SetHeader("cookie", oldCookies)
	}

	resp, err := request.Get(url)

	if err != nil {
		r.logger.Err(err).Msg("RequestHtmlService.Do.Error")
		return nil, err
	}

	cookie := resp.Header().Get("set-cookie")

	if cookie != "" {
		r.logger.Info().Str("url", url).Str("cookie", cookie).Msg("RequestHtmlService.SetCookie")
		r.cookieStorage.SetCookie(url, cookie)
	}

	r.logger.Info().
		Str("url", url).
		Str("code", resp.Status()).
		Msg("ScrapHtml.Success")

	return &entity.HtmlPage{
		Html:       resp.String(),
		StatusCode: resp.StatusCode(),
	}, nil

}
