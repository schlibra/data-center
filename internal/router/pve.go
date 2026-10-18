package router

import (
	"data-center/internal/handler/pve"

	"github.com/gin-gonic/gin"
)

func initPveNodeRouter(router *gin.RouterGroup) {
	node := router.Group("/node")
	node.GET("/", pve.ListPveNodeHandler)
}

func initPveQemuRouter(router *gin.RouterGroup) {
	qemu := router.Group("/qemu")
	qemu.GET("/:node", pve.ListPveQemuHandler)
}

func initPveQemuStatusRouter(router *gin.RouterGroup) {
	status := router.Group("/qemu-status")
	status.GET("/:node/:vmid", pve.GetPveQemuStatusHandler)
	status.POST("/:node/:vmid", pve.SetPveQemuStatusHandler)
}

func initPveLxcRouter(router *gin.RouterGroup) {
	lxc := router.Group("/lxc")
	lxc.GET("/:node", pve.ListPveLxcHandler)
}
func initPveLxcStatusRouter(router *gin.RouterGroup) {
	status := router.Group("/lxc-status")
	status.GET("/:node/:vmid", pve.GetPveLxcStatusHandler)
	status.POST("/:node/:vmid", pve.SetPveLxcStatusHandler)
}
func initPveRouter(router *gin.RouterGroup) {
	pveRouter := router.Group("/pve")
	initPveNodeRouter(pveRouter)
	initPveQemuRouter(pveRouter)
	initPveQemuStatusRouter(pveRouter)
	initPveLxcRouter(pveRouter)
	initPveLxcStatusRouter(pveRouter)
}
