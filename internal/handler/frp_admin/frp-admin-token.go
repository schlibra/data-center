package frp_admin

import (
	frpAdminModel "data-center/internal/models/request/frp_admin"
	"data-center/internal/service/frp_admin"

	"github.com/gin-gonic/gin"
)

// ListFrpAdminTokenHandler 获取 FRP 令牌列表（管理员）
// @Summary 获取 FRP 令牌列表（管理员）
// @Tags FRP管理-令牌
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.FrpTokenTable}
// @Router /frp/admin/token/ [get]
func ListFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	frp_admin.ListFrpAdminTokenService(c)
}

// GetFrpAdminTokenHandler 获取 FRP 令牌详情（管理员）
// @Summary 获取 FRP 令牌详情（管理员）
// @Tags FRP管理-令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response{data=models.FrpTokenTable}
// @Router /frp/admin/token/{id} [get]
func GetFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	frp_admin.GetFrpAdminTokenService(c, id)
}

// CreateFrpAdminTokenHandler 创建 FRP 令牌（管理员）
// @Summary 创建 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body frp_admin.CreateFrpAdminToken true "令牌信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/ [post]
func CreateFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	req := frpAdminModel.CreateFrpAdminToken{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	frp_admin.CreateFrpAdminTokenService(c, req)
}

// UpdateFrpAdminTokenHandler 更新 FRP 令牌（管理员）
// @Summary 更新 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Param request body frp_admin.UpdateFrpAdminToken true "令牌信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/{id} [put]
func UpdateFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	req := frpAdminModel.UpdateFrpAdminToken{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	frp_admin.UpdateFrpAdminTokenService(c, id, req)
}

// DeleteFrpAdminTokenHandler 删除 FRP 令牌（管理员）
// @Summary 删除 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/{id} [delete]
func DeleteFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	frp_admin.DeleteFrpAdminTokenService(c, id)
}

// GenerateFrpAdminTokenHandler 重新生成 FRP 令牌（管理员）
// @Summary 重新生成 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/{id} [post]
func GenerateFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	frp_admin.GenerateFrpAdminTokenService(c, id)
}
