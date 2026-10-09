package router

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func initFrpAdminRuleRouter(router *gin.RouterGroup) {
	rule := router.Group("/rule")
	rule.GET("/", handler.ListFrpAdminRuleHandler)
	rule.GET("/:id", handler.GetFrpAdminRuleHandler)
	rule.POST("/", handler.CreateFrpAdminRuleHandler)
	rule.PUT("/:id", handler.UpdateFrpAdminRuleHandler)
	rule.DELETE("/:id", handler.DeleteFrpAdminRuleHandler)
}

func initFrpAdminTokenRouter(router *gin.RouterGroup) {
	token := router.Group("/token")
	token.GET("/", handler.ListFrpAdminTokenHandler)
	token.GET("/:id", handler.GetFrpAdminTokenHandler)
	token.PUT("/:id", handler.UpdateFrpAdminTokenHandler)
	token.DELETE("/:id", handler.DeleteFrpAdminTokenHandler)
	token.POST("/", handler.CreateFrpAdminTokenHandler)
	token.POST("/:id", handler.GenerateFrpAdminTokenHandler)
}
func initFrpAdminApiRouter(router *gin.RouterGroup) {
	api := router.Group("/api")
	api.GET("/proxy", handler.ProxiesFrpAdminApiHandler)
	api.GET("/client", handler.ClientsFrpAdminApiHandler)
}

func initFrpAdminRouter(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	initFrpAdminRuleRouter(admin)
	initFrpAdminTokenRouter(admin)
	initFrpAdminApiRouter(admin)
}

func initFrpApiRouter(router *gin.RouterGroup) {
	api := router.Group("/api")
	api.POST("/login", handler.LoginFrpApiHandler)
	api.POST("/proxy", handler.ProxyFrpApiHandler)
	api.GET("/proxy", handler.ProxiesFrpApiHandler)
	api.GET("/client", handler.ClientsFrpApiHandler)
	api.GET("/info", handler.InfoFrpApiHandler)
}

func initFrpTokenRouter(router *gin.RouterGroup) {
	token := router.Group("/token")
	token.POST("/", handler.CreateFrpTokenHandler)
	token.GET("/", handler.ListFrpTokenHandler)
	token.GET("/:id", handler.GetFrpTokenHandler)
	token.PUT("/:id", handler.UpdateFrpTokenHandler)
	token.DELETE("/:id", handler.DeleteFrpTokenHandler)
	token.POST("/:id", handler.GenerateFrpTokenHandler)
}

func initFrpRuleRouter(router *gin.RouterGroup) {
	rule := router.Group("/rule")
	rule.POST("/", handler.CreateFrpRuleHandler)
	rule.GET("/", handler.ListFrpRuleHandler)
	rule.GET("/:id", handler.GetFrpRuleHandler)
	rule.PUT("/:id", handler.UpdateFrpRuleHandler)
	rule.DELETE("/:id", handler.DeleteFrpRuleHandler)
}

func initFrpRouter(router *gin.RouterGroup) {
	frp := router.Group("/frp")
	initFrpTokenRouter(frp)
	initFrpRuleRouter(frp)
	initFrpApiRouter(frp)
	initFrpAdminRouter(frp)
}
