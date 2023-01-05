package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/millirud/go-scrapper/internal/di"
	"github.com/millirud/go-scrapper/internal/service/request_html"
	"github.com/rs/zerolog"
)

type htmlRequestRouter struct {
	logger             *zerolog.Logger
	HtmlRequestService request_html.RequestHtmlService
}

func newHtmlRequestRouter(handler *gin.RouterGroup, Di *di.DI) {
	r := &htmlRequestRouter{
		logger:             Di.Logger,
		HtmlRequestService: *Di.RequestHtmlService,
	}

	handler.POST("/scrap-html", r.ScrapHtml)
}

// Booking contains binded and validated data.
type ScrapHtml struct {
	Url string `form:"url" binding:"required"`
}

func (r *htmlRequestRouter) ScrapHtml(c *gin.Context) {
	var payload ScrapHtml

	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.HtmlRequestService.Do(payload.Url)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
