package exception

import (
	"URLShortener/exception/custom_error"
	"URLShortener/helper"
	"URLShortener/model/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {
	if linkNotFoundError(c, recovered) {
		return
	}

	if invalidUrlError(c, recovered) {
		return
	}

	internalServerError(c, recovered)
}

func linkNotFoundError(c *gin.Context, recovered interface{}) bool {
	_, ok := recovered.(custom_error.LinkNotFoundError)

	if ok {
		//webRes := web.WebResponse{
		//	Code:   http.StatusNotFound,
		//	Status: "The Page you search is not found",
		//	Data:   recover.Error,
		//}
		//
		//c.JSON(http.StatusNotFound, webRes)
		clientPath := helper.Getenv("CLIENT_PATH", "https://johaneswiku.com/urlshort")
		c.Redirect(http.StatusMovedPermanently, clientPath)
		return true
	}

	return false
}

func invalidUrlError(c *gin.Context, recovered interface{}) bool {
	recover, ok := recovered.(custom_error.InvalidUrl)

	if ok {
		webRes := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Url !",
			Data:   recover.Error,
		}

		c.JSON(http.StatusBadRequest, webRes)
		return true
	}

	return false
}

func internalServerError(c *gin.Context, recovered interface{}) {
	webRes := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   recovered,
	}

	c.JSON(http.StatusInternalServerError, webRes)
}
