package router

import (
	"data-center/internal/handler/openvpn"

	"github.com/gin-gonic/gin"
)

func initOpenVPNRouter(router *gin.RouterGroup) {
	openvpnRouter := router.Group("/openvpn")
	initOpenVPNUserRouter(openvpnRouter)
	initOpenVPNGroupRouter(openvpnRouter)
	initOpenVPNClientRouter(openvpnRouter)
}
func initOpenVPNClientRouter(router *gin.RouterGroup) {
	client := router.Group("/client")
	client.GET("/", openvpn.ListOpenVPNClientHandler)
	client.GET("/:id", openvpn.GetOpenVPNClientHandler)
	client.DELETE("/:id", openvpn.KickOpenVPNClientHandler)
}
func initOpenVPNGroupRouter(router *gin.RouterGroup) {
	group := router.Group("/group")
	group.GET("/", openvpn.ListOpenVPNGroupHandler)
	group.GET("/:id", openvpn.GetOpenVPNGroupHandler)
	group.POST("/", openvpn.CreateOpenVPNGroupHandler)
	group.PUT("/:id", openvpn.UpdateOpenVPNGroupHandler)
	group.DELETE("/:id", openvpn.DeleteOpenVPNGroupHandler)
}
func initOpenVPNUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("/", openvpn.ListOpenVPNUserHandler)
	user.GET("/:id", openvpn.GetOpenVPNUserHandler)
	user.POST("/", openvpn.CreateOpenVPNUserHandler)
	user.PUT("/:id", openvpn.UpdateOpenVPNUserHandler)
	user.DELETE("/:id", openvpn.DeleteOpenVPNUserHandler)
}
