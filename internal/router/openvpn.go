package router

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func initOpenVPNRouter(router *gin.RouterGroup) {
	openvpn := router.Group("/openvpn")
	initOpenVPNUserRouter(openvpn)
	initOpenVPNGroupRouter(openvpn)
	initOpenVPNClientRouter(openvpn)
}
func initOpenVPNClientRouter(router *gin.RouterGroup) {
	client := router.Group("/client")
	client.GET("/", handler.ListOpenVPNClientHandler)
	client.GET("/:id", handler.GetOpenVPNClientHandler)
	client.DELETE("/:id", handler.KickOpenVPNClientHandler)
}
func initOpenVPNGroupRouter(router *gin.RouterGroup) {
	group := router.Group("/group")
	group.GET("/", handler.ListOpenVPNGroupHandler)
	group.GET("/:id", handler.GetOpenVPNGroupHandler)
	group.POST("/", handler.CreateOpenVPNGroupHandler)
	group.PUT("/:id", handler.UpdateOpenVPNGroupHandler)
	group.DELETE("/:id", handler.DeleteOpenVPNGroupHandler)
}
func initOpenVPNUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("/", handler.ListOpenVPNUserHandler)
	user.GET("/:id", handler.GetOpenVPNUserHandler)
	user.POST("/", handler.CreateOpenVPNUserHandler)
	user.PUT("/:id", handler.UpdateOpenVPNUserHandler)
	user.DELETE("/:id", handler.DeleteOpenVPNUserHandler)
}
