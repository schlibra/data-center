package frp_admin

import (
	"data-center/internal/service/frp_admin"

	"github.com/gin-gonic/gin"
)

// ProxiesFrpAdminApiHandler 获取全部代理列表（管理员）
// @Summary 获取全部代理列表（管理员）
// @Tags FRP管理-接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]frp.FrpApiProxiesDataItem}
// @Router /frp/admin/api/proxy [get]
func ProxiesFrpAdminApiHandler(c *gin.Context) {
	checkAdmin(c)
	frp_admin.ProxiesFrpAdminApiService(c)
}

// ClientsFrpAdminApiHandler 获取全部客户端列表（管理员）
// @Summary 获取全部客户端列表（管理员）
// @Tags FRP管理-接口
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]frp.FrpApiClientsDataItem}
// @Router /frp/admin/api/client [get]
func ClientsFrpAdminApiHandler(c *gin.Context) {
	checkAdmin(c)
	frp_admin.ClientsFrpAdminApiService(c)
}
