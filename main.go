package main

import (
	"URLShortener/helper"
	"URLShortener/injector"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load()
	helper.PanicIfError(err)

	r := injector.InitializeServer()
	if err := r.Run("127.0.0.1:8092"); err == nil {
		msg := fmt.Sprintf("Failed to start server: Error %v", err)
		panic(msg)
	}
}
