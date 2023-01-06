package cookie_local_storage

import (
	"net/url"
	"sync"

	"github.com/rs/zerolog"
)

func New(logger *zerolog.Logger) *CookieLocalStorage {
	return &CookieLocalStorage{
		logger: logger,
		store:  make(map[string]string),
	}
}

type CookieLocalStorage struct {
	sync.RWMutex
	logger *zerolog.Logger
	store  map[string]string
}

func (c *CookieLocalStorage) SetCookie(uri string, cookie string) error {
	c.logger.Info().Str("uri", uri).Str("cookie", cookie).Msg("SetCookie")

	c.Lock()
	defer c.Unlock()

	host, err := c.hostFromUrl(uri)

	if err != nil {
		c.logger.Err(err).Str("uri", uri).Str("host", host).Msg("SetCookie.Err")
		return err
	}

	c.store[host] = cookie

	return nil
}

func (c *CookieLocalStorage) GetCookie(uri string) (cookie string, err error) {
	c.logger.Info().Str("uri", uri).Msg("GetCookie")

	c.RLock()
	defer c.RUnlock()

	host, err := c.hostFromUrl(uri)

	if err != nil {
		c.logger.Err(err).Str("uri", uri).Str("host", host).Msg("GetCookie.Err")
		return "", err
	}

	return c.store[host], nil
}

func (c *CookieLocalStorage) hostFromUrl(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	return u.Hostname(), nil
}
