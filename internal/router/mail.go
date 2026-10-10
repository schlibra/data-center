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

func initMailRouter(router *gin.RouterGroup) {
	mailRouter := router.Group("/mail")
	initMailDomainRouter(mailRouter)
}
