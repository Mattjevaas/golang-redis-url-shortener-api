package store_controller

import (
	"URLShortener/helper"
	"URLShortener/model/web"
	"URLShortener/service/store_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StoreControllerImpl struct {
	StoreService store_service.StoreService
}

func NewStoreController(storeService store_service.StoreService) StoreContoller {
	return &StoreControllerImpl{
		StoreService: storeService,
	}
}

func (s *StoreControllerImpl) SaveURLMapping(c *gin.Context) {
	webReq := web.WebRequest{}
	err := c.BindJSON(&webReq)
	helper.PanicIfError(err)

	shortUrl := s.StoreService.SaveURLMapping(c, webReq)

	webRes := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   map[string]string{"short_url": shortUrl},
	}

	c.JSON(http.StatusOK, webRes)
}

func (s *StoreControllerImpl) RetrieveInitialURL(c *gin.Context) {
	param := c.Param("shortUrl")
	if param == "" {
		panic("ShortURL Cannot be Empty")
	}

	oriUrl := s.StoreService.RetrieveInitialURL(c, param)

	webRes := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   map[string]string{"original_url": oriUrl},
	}

	c.JSON(http.StatusOK, webRes)
}
