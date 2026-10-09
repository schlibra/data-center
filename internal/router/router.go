package router

import (
	_ "data-center/docs"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/api")
	initUserRouter(api)
	initFrpRouter(api)
	initAdminRouter(api)
	initOpenVPNRouter(api)
	initConfigRouter(api)
}
