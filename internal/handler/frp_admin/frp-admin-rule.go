package frp_admin

import (
	frpAdminModel "data-center/internal/models/request/frp_admin"
	"data-center/internal/service/frp_admin"

	"github.com/gin-gonic/gin"
)

// ListFrpAdminRuleHandler 获取 FRP 规则列表（管理员）
// @Summary 获取 FRP 规则列表（管理员）
// @Tags FRP管理-规则
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.FrpRuleTable}
// @Router /frp/admin/rule/ [get]
func ListFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	frp_admin.ListFrpAdminRuleService(c)
}

// GetFrpAdminRuleHandler 获取 FRP 规则详情（管理员）
// @Summary 获取 FRP 规则详情（管理员）
// @Tags FRP管理-规则
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Success 200 {object} response.Response{data=models.FrpRuleTable}
// @Router /frp/admin/rule/{id} [get]
func GetFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	frp_admin.GetFrpAdminRuleService(c, id)
}

// CreateFrpAdminRuleHandler 创建 FRP 规则（管理员）
// @Summary 创建 FRP 规则（管理员）
// @Tags FRP管理-规则
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body frp_admin.CreateFrpAdminRule true "规则信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/rule/ [post]
func CreateFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	req := frpAdminModel.CreateFrpAdminRule{}.Parse(c)
	portCheck(c, req.Min, req.Max)
	frp_admin.CreateFrpAdminRuleService(c, req)
}

// UpdateFrpAdminRuleHandler 更新 FRP 规则（管理员）
// @Summary 更新 FRP 规则（管理员）
// @Tags FRP管理-规则
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Param request body frp_admin.UpdateFrpAdminRule true "规则信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/rule/{id} [put]
func UpdateFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	req := frpAdminModel.UpdateFrpAdminRule{}.Parse(c)
	portCheck(c, req.Min, req.Max)
	frp_admin.UpdateFrpAdminRuleService(c, req, id)
}

// DeleteFrpAdminRuleHandler 删除 FRP 规则（管理员）
// @Summary 删除 FRP 规则（管理员）
// @Tags FRP管理-规则
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Success 200 {object} response.Response
// @Router /frp/admin/rule/{id} [delete]
func DeleteFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	frp_admin.DeleteFrpAdminRuleService(c, id)
}
