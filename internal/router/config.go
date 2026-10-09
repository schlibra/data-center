package router

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func initConfigRouter(router *gin.RouterGroup) {
	config := router.Group("/config")
	config.GET("/:key", handler.GetConfigHandler)
}
