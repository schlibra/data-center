package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

func LoginFrpApiHandler(c *gin.Context) {
	var req request.FrpApiLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.LoginFrpApiService(c, req)
}
func ProxyFrpApiHandler(c *gin.Context) {
	var req request.FrpApiProxy
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.ProxyFrpApiService(c, req)
}
func ProxiesFrpApiHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.api.proxy")
	service.ProxiesFrpApiService(c, row)
}
func ClientsFrpApiHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.api.clients")
	service.ClientFrpApiService(c, row)
}
