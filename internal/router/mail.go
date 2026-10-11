package router

import (
	"data-center/internal/handler/mail"

	"github.com/gin-gonic/gin"
)

func initMailDomainRouter(router *gin.RouterGroup) {
	domain := router.Group("/domain")
	domain.GET("/", mail.ListMailDomainHandler)
	domain.GET("/:id", mail.GetMailDomainHandler)
	domain.POST("/", mail.CreateMailDomainHandler)
	domain.PUT("/:id", mail.UpdateMailDomainHandler)
	domain.DELETE("/:id", mail.DeleteMailDomainHandler)
}

func initMailListRouter(router *gin.RouterGroup) {
	list := router.Group("/list")
	list.GET("/:domain", mail.ListMailListHandler)
	list.GET("/:domain/:record", mail.GetMailListHandler)
	list.POST("/:domain", mail.CreateMailListHandler)
	list.PUT("/:domain/:record", mail.UpdateMailListHandler)
	list.DELETE("/:domain/:record", mail.DeleteMailListHandler)
}

func initMailRouter(router *gin.RouterGroup) {
	mailRouter := router.Group("/mail")
	initMailDomainRouter(mailRouter)
	initMailListRouter(mailRouter)
}
