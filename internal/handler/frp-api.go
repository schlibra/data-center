package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// LoginFrpApiHandler FRP 客户端登录回调
// @Summary FRP 客户端登录回调（frps webhook）
// @Tags FRP接口
// @Accept application/json
// @Produce application/json
// @Param request body request.FrpApiLogin true "登录回调参数"
// @Success 200 {object} response.Response
// @Router /frp/api/login [post]
func LoginFrpApiHandler(c *gin.Context) {
	var req request.FrpApiLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.LoginFrpApiService(c, req)
}

// ProxyFrpApiHandler FRP 创建代理回调
// @Summary FRP 创建代理回调（frps webhook）
// @Tags FRP接口
// @Accept application/json
// @Produce application/json
// @Param request body request.FrpApiProxy true "代理回调参数"
// @Success 200 {object} response.Response
// @Router /frp/api/proxy [post]
func ProxyFrpApiHandler(c *gin.Context) {
	var req request.FrpApiProxy
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.ProxyFrpApiService(c, req)
}

// ProxiesFrpApiHandler 获取当前用户代理列表
// @Summary 获取当前用户代理列表
// @Tags FRP接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]request.FrpApiProxiesDataItem}
// @Router /frp/api/proxy [get]
func ProxiesFrpApiHandler(c *gin.Context) {
	row := checkPermission(c, "frp.api.proxy")
	service.ProxiesFrpApiService(c, row)
}

// ClientsFrpApiHandler 获取当前用户客户端列表
// @Summary 获取当前用户客户端列表
// @Tags FRP接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]request.FrpApiClientsDataItem}
// @Router /frp/api/client [get]
func ClientsFrpApiHandler(c *gin.Context) {
	row := checkPermission(c, "frp.api.client")
	service.ClientFrpApiService(c, row)
}
