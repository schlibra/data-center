package router

import (
	"data-center/internal/handler/frp"
	"data-center/internal/handler/frp_admin"

	"github.com/gin-gonic/gin"
)

func initFrpAdminRuleRouter(router *gin.RouterGroup) {
	rule := router.Group("/rule")
	rule.GET("/", frp_admin.ListFrpAdminRuleHandler)
	rule.GET("/:id", frp_admin.GetFrpAdminRuleHandler)
	rule.POST("/", frp_admin.CreateFrpAdminRuleHandler)
	rule.PUT("/:id", frp_admin.UpdateFrpAdminRuleHandler)
	rule.DELETE("/:id", frp_admin.DeleteFrpAdminRuleHandler)
}

func initFrpAdminTokenRouter(router *gin.RouterGroup) {
	token := router.Group("/token")
	token.GET("/", frp_admin.ListFrpAdminTokenHandler)
	token.GET("/:id", frp_admin.GetFrpAdminTokenHandler)
	token.PUT("/:id", frp_admin.UpdateFrpAdminTokenHandler)
	token.DELETE("/:id", frp_admin.DeleteFrpAdminTokenHandler)
	token.POST("/", frp_admin.CreateFrpAdminTokenHandler)
	token.POST("/:id", frp_admin.GenerateFrpAdminTokenHandler)
}
func initFrpAdminApiRouter(router *gin.RouterGroup) {
	api := router.Group("/api")
	api.GET("/proxy", frp_admin.ProxiesFrpAdminApiHandler)
	api.DELETE("/proxy", frp_admin.ClearOfflineFrpAdminApiHandler)
	api.GET("/client", frp_admin.ClientsFrpAdminApiHandler)
}

func initFrpAdminRouter(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	initFrpAdminRuleRouter(admin)
	initFrpAdminTokenRouter(admin)
	initFrpAdminApiRouter(admin)
}

func initFrpApiRouter(router *gin.RouterGroup) {
	api := router.Group("/api")
	api.POST("/login", frp.LoginFrpApiHandler)
	api.POST("/proxy", frp.ProxyFrpApiHandler)
	api.GET("/proxy", frp.ProxiesFrpApiHandler)
	api.GET("/client", frp.ClientsFrpApiHandler)
	api.GET("/info", frp.InfoFrpApiHandler)
}

func initFrpTokenRouter(router *gin.RouterGroup) {
	token := router.Group("/token")
	token.POST("/", frp.CreateFrpTokenHandler)
	token.GET("/", frp.ListFrpTokenHandler)
	token.GET("/:id", frp.GetFrpTokenHandler)
	token.PUT("/:id", frp.UpdateFrpTokenHandler)
	token.DELETE("/:id", frp.DeleteFrpTokenHandler)
	token.POST("/:id", frp.GenerateFrpTokenHandler)
}

func initFrpRuleRouter(router *gin.RouterGroup) {
	rule := router.Group("/rule")
	rule.POST("/", frp.CreateFrpRuleHandler)
	rule.GET("/", frp.ListFrpRuleHandler)
	rule.GET("/:id", frp.GetFrpRuleHandler)
	rule.PUT("/:id", frp.UpdateFrpRuleHandler)
	rule.DELETE("/:id", frp.DeleteFrpRuleHandler)
}

func initFrpRouter(router *gin.RouterGroup) {
	frp := router.Group("/frp")
	initFrpTokenRouter(frp)
	initFrpRuleRouter(frp)
	initFrpApiRouter(frp)
	initFrpAdminRouter(frp)
}
