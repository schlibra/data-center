package handler

import (
	"data-center/internal/models/request"
	"data-center/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// ProxiesFrpAdminApiHandler 获取全部代理列表（管理员）
// @Summary 获取全部代理列表（管理员）
// @Tags FRP管理-接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]request.FrpApiProxiesDataItem}
// @Router /frp/admin/api/proxy [get]
func ProxiesFrpAdminApiHandler(c *gin.Context) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	client := resty.New()
	var proxyResult request.FrpApiProxies
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&proxyResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/proxies?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", proxyResult.Data.Items)
}

// ClientsFrpAdminApiHandler 获取全部客户端列表（管理员）
// @Summary 获取全部客户端列表（管理员）
// @Tags FRP管理-接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]request.FrpApiClientsDataItem}
// @Router /frp/admin/api/client [get]
func ClientsFrpAdminApiHandler(c *gin.Context) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	client := resty.New()
	var clientResult request.FrpApiClients
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&clientResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/clients?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", clientResult.Data.Items)
}
