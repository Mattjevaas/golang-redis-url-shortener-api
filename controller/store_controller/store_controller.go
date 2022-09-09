package store_controller

import (
	"github.com/gin-gonic/gin"
)

type StoreContoller interface {
	SaveURLMapping(c *gin.Context)
	RetrieveInitialURL(c *gin.Context)
}
