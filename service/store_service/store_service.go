package store_service

import (
	"URLShortener/model/web"
	"github.com/gin-gonic/gin"
)

type StoreService interface {
	SaveURLMapping(c *gin.Context, request web.WebRequest) string
	RetrieveInitialURL(c *gin.Context, shortUrl string) string
}
