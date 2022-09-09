package app

import (
	"URLShortener/controller/store_controller"
	"URLShortener/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewRouter(
	redis *redis.Client,
	storeController store_controller.StoreContoller,
) *gin.Engine {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		pong, err := redis.Ping(c).Result()
		helper.PanicIfError(err)
		fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

		c.Set("redis", redis)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to GO URL Shortener!"})
	})

	r.POST("/save", storeController.SaveURLMapping)
	r.GET("/retrieve/:shortUrl", storeController.RetrieveInitialURL)

	return r
}
