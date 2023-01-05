package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/millirud/go-scrapper/internal/di"
)

func NewRouter(handler *gin.Engine, Di *di.DI) {

	handler.GET("/healthz", newHealthz())

	// Routers
	h := handler.Group("/v1")
	{
		newHtmlRequestRouter(h, Di)
	}
}

func newHealthz() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
