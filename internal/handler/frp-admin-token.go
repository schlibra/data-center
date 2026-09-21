package handler

import "github.com/gin-gonic/gin"

// ListFrpAdminTokenHandler 获取 FRP 令牌列表（管理员）
// @Summary 获取 FRP 令牌列表（管理员）
// @Tags FRP管理-令牌
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.FrpTokenTable}
// @Router /frp/admin/token/ [get]
func ListFrpAdminTokenHandler(c *gin.Context) {

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

}

// CreateFrpAdminTokenHandler 创建 FRP 令牌（管理员）
// @Summary 创建 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/admin/token/ [post]
func CreateFrpAdminTokenHandler(c *gin.Context) {

}

// UpdateFrpAdminTokenHandler 更新 FRP 令牌（管理员）
// @Summary 更新 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/{id} [put]
func UpdateFrpAdminTokenHandler(c *gin.Context) {

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

}
