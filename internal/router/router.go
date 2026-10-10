package router

import (
	_ "data-center/docs"

	"github.com/gin-gonic/gin"
)

var (
	Version   string
	Commit    string
	BuildTime string
)

func InitRouter(router *gin.Engine, version string, commit string, buildTime string) {
	Version = version
	Commit = commit
	BuildTime = buildTime
	api := router.Group("/api")
	initUserRouter(api)
	initFrpRouter(api)
	initAdminRouter(api)
	initOpenVPNRouter(api)
	initConfigRouter(api)
}
