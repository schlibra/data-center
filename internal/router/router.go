package router

import (
	_ "data-center/docs"
	"data-center/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST("/login", handler.UserLoginHandler)
	user.PUT("/login", handler.UserLoginKeyHandler)
	user.POST("/register", handler.UserRegisterHandler)
	user.PUT("/register", handler.UserRegisterKeyHandler)
	user.POST("/logout", handler.UserLogoutHandler)
	user.GET("/", handler.UserInfoHandler)
	user.PUT("/", handler.UserUpdateHandler)
	user.PATCH("/", handler.UserPasswordHandler)
	user.POST("/api", handler.UserApiKeyHandler)
}

func InitRouter(router *gin.Engine) {
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	initUserRouter(api)
	initFrpRouter(api)
	initAdminRouter(api)
}
