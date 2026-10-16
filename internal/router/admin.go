package router

import (
	"data-center/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

func initAdminUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("/", admin.ListAdminUserHandler)
	user.GET("/:id", admin.GetAdminUserHandler)
	user.POST("/", admin.CreateAdminUserHandler)
	user.PUT("/:id", admin.UpdateAdminUserHandler)
	user.PATCH("/:id", admin.PasswordAdminUserHandler)
	user.DELETE("/:id", admin.DeleteAdminUserHandler)
}
func initAdminGroupRouter(router *gin.RouterGroup) {
	group := router.Group("/group")
	group.GET("/", admin.ListAdminGroupHandler)
	group.GET("/:id", admin.GetAdminGroupHandler)
	group.POST("/", admin.CreateAdminGroupHandler)
	group.PUT("/:id", admin.UpdateAdminGroupHandler)
	group.DELETE("/:id", admin.DeleteAdminGroupHandler)
}
func initAdminPermissionRouter(router *gin.RouterGroup) {
	permission := router.Group("/permission")
	permission.GET("/", admin.ListAdminPermissionsHandler)
	permission.POST("/", admin.CreateAdminPermissionHandler)
	permission.GET("/:id", admin.GetAdminPermissionHandler)
	permission.PUT("/:id", admin.UpdateAdminPermissionHandler)
	permission.DELETE("/:id", admin.DeleteAdminPermissionHandler)
}
func initAdminSettingsRouter(router *gin.RouterGroup) {
	settings := router.Group("/settings")
	settings.GET("/", admin.ListAdminSettingsHandler)
	settings.GET("/:id", admin.GetAdminSettingsHandler)
	settings.POST("/", admin.CreateAdminSettingsHandler)
	settings.PUT("/", admin.SetAdminSettingsHandler)
	settings.PUT("/:id", admin.UpdateAdminSettingsHandler)
	settings.DELETE("/:id", admin.DeleteAdminSettingsHandler)
}

func initAdminVersionRouter(router *gin.RouterGroup) {
	version := router.Group("/version")
	version.GET("/", func(c *gin.Context) {
		admin.GetAdminVersionHandler(c, Version, Commit, BuildTime)
	})
	version.POST("/", admin.UpgradeAdminVersionHandler)
}
func initAdminBenchmarkRouter(router *gin.RouterGroup) {
	benchmark := router.Group("/benchmark")
	benchmark.Any("/", admin.DoAdminBenchmarkHandler)
}
func initAdminRouter(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")
	initAdminUserRouter(adminRouter)
	initAdminGroupRouter(adminRouter)
	initAdminPermissionRouter(adminRouter)
	initAdminSettingsRouter(adminRouter)
	initAdminVersionRouter(adminRouter)
	initAdminBenchmarkRouter(adminRouter)
}
