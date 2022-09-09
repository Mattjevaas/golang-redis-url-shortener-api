package store_repository

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type StoreRepository interface {
	SaveURLMapping(c *gin.Context, redis *redis.Client, shortUrl string, originalUrl string)
	RetrieveInitialURL(c *gin.Context, redis *redis.Client, shortUrl string) string
}
