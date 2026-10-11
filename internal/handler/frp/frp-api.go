package frp

import (
	frpModel "data-center/internal/models/request/frp"
	"data-center/internal/service/frp"

	"github.com/gin-gonic/gin"
)

// LoginFrpApiHandler FRP 客户端登录回调
// @Summary FRP 客户端登录回调（frps webhook）
// @Tags FRP接口
// @Accept application/json
// @Produce application/json
// @Param request body frp.LoginFrpApi true "登录回调参数"
// @Success 200 {object} response.Response
// @Router /frp/api/login [post]
func LoginFrpApiHandler(c *gin.Context) {
	req := frpModel.LoginFrpApi{}.Parse(c)
	frp.LoginFrpApiService(c, req)
}

// ProxyFrpApiHandler FRP 创建代理回调
// @Summary FRP 创建代理回调（frps webhook）
// @Tags FRP接口
// @Accept application/json
// @Produce application/json
// @Param request body frp.ProxyFrpApi true "代理回调参数"
// @Success 200 {object} response.Response
// @Router /frp/api/proxy [post]
func ProxyFrpApiHandler(c *gin.Context) {
	req := frpModel.ProxyFrpApi{}.Parse(c)
	frp.ProxyFrpApiService(c, req)
}

// ProxiesFrpApiHandler 获取当前用户代理列表
// @Summary 获取当前用户代理列表
// @Tags FRP接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]frp.ProxiesFrpApiDataItem}
// @Router /frp/api/proxy [get]
func ProxiesFrpApiHandler(c *gin.Context) {
	row := checkPermission(c, "frp.api.proxy")
	frp.ProxiesFrpApiService(c, row)
}

// ClientsFrpApiHandler 获取当前用户客户端列表
// @Summary 获取当前用户客户端列表
// @Tags FRP接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]frp.ClientsFrpApiDataItem}
// @Router /frp/api/client [get]
func ClientsFrpApiHandler(c *gin.Context) {
	row := checkPermission(c, "frp.api.client")
	frp.ClientFrpApiService(c, row)
}

// InfoFrpApiHandler 获取Frp信息
// @Summary 获取Frp信息
// @Tags FRP接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/api/info [get]
func InfoFrpApiHandler(c *gin.Context) {
	checkPermission(c, "frp.api.info")
	frp.InfoFrpApiService(c)
}
