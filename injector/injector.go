//go:build wireinject
// +build wireinject

package injector

import (
	"URLShortener/app"
	"URLShortener/controller/store_controller"
	"URLShortener/repository/store_repository"
	"URLShortener/service/store_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitializeServer() *gin.Engine {
	wire.Build(
		app.NewRouter,
		app.NewRedisClient,
		validator.New,
		store_controller.NewStoreController,
		store_service.NewStoreService,
		store_repository.NewStoreRepository,
	)

	return nil
}
