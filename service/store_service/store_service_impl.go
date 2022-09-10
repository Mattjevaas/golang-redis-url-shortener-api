package store_service

import (
	"URLShortener/exception/custom_error"
	"URLShortener/helper"
	"URLShortener/model/web"
	"URLShortener/repository/store_repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
)

type StoreServiceImpl struct {
	StoreRepository store_repository.StoreRepository
	Validator       *validator.Validate
}

func NewStoreService(storeRepository store_repository.StoreRepository, validator *validator.Validate) StoreService {
	return &StoreServiceImpl{
		StoreRepository: storeRepository,
		Validator:       validator,
	}
}

func (s *StoreServiceImpl) SaveURLMapping(c *gin.Context, request web.WebRequest) string {
	errValiadate := s.Validator.Struct(request)
	if errValiadate != nil {
		panic(custom_error.NewInvalidUrl(errValiadate.Error()))
	}

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
