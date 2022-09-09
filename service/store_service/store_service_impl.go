package store_service

import (
	"URLShortener/helper"
	"URLShortener/model/web"
	"URLShortener/repository/store_repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type StoreServiceImpl struct {
	StoreRepository store_repository.StoreRepository
}

func NewStoreService(storeRepository store_repository.StoreRepository) StoreService {
	return &StoreServiceImpl{
		StoreRepository: storeRepository,
	}
}

func (s *StoreServiceImpl) SaveURLMapping(c *gin.Context, request web.WebRequest) string {
	redis, err := c.MustGet("redis").(*redis.Client)

	if !err {
		helper.PanicIfError(errors.New("cannot get redis"))
	}

	shortUrl := helper.GenerateShortLink(request.OriginalUrl)
	s.StoreRepository.SaveURLMapping(c, redis, shortUrl, request.OriginalUrl)

	return shortUrl
}

func (s *StoreServiceImpl) RetrieveInitialURL(c *gin.Context, shortUrl string) string {
	redis, err := c.MustGet("redis").(*redis.Client)

	if !err {
		helper.PanicIfError(errors.New("cannot get redis"))
	}

	res := s.StoreRepository.RetrieveInitialURL(c, redis, shortUrl)
	return res
}
