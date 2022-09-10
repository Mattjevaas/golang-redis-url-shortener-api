package store_repository

import (
	"URLShortener/exception/custom_error"
	"URLShortener/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type StoreRepositoryImpl struct{}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (s *StoreRepositoryImpl) SaveURLMapping(c *gin.Context, redis *redis.Client, shortUrl string, originalUrl string) {
	ttl := helper.Getenv("TTL", "1")
	ttlInt, err := strconv.Atoi(ttl)
	helper.PanicIfError(err)

	redisErr := redis.Set(c, shortUrl, originalUrl, time.Duration(ttlInt)*time.Hour).Err()
	helper.PanicIfError(redisErr)
}

func (s *StoreRepositoryImpl) RetrieveInitialURL(c *gin.Context, redis *redis.Client, shortUrl string) string {
	result, redisErr := redis.Get(c, shortUrl).Result()

	if redisErr != nil {
		panic(custom_error.NewLinkNotFound(redisErr.Error()))
	}

	return result
}
