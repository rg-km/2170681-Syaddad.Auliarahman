package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {

	route := gin.Default()

	route.GET("/:url", urlHandler.Get)
	route.POST("/", urlHandler.Create)
	route.POST("/custom", urlHandler.CreateCustom)

	return route
	// TODO: replace this
}
