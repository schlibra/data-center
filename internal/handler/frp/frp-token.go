package frp

import (
	frpModel "data-center/internal/models/request/frp"
	"data-center/internal/service/frp"

	"github.com/gin-gonic/gin"
)

// CreateFrpTokenHandler 创建 FRP 令牌
// @Summary 创建 FRP 令牌
// @Tags FRP令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body frp.CreateFrpToken true "令牌参数"
// @Success 200 {object} response.Response
// @Router /frp/token/ [post]
func CreateFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.create")
	req := frpModel.CreateFrpToken{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	frp.CreateFrpTokenService(c, req, row)
}

// ListFrpTokenHandler 获取 FRP 令牌列表
// @Summary 获取 FRP 令牌列表
// @Tags FRP令牌
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/token/ [get]
func ListFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.get")
	frp.ListFrpTokenService(c, row)
}

// GetFrpTokenHandler 获取 FRP 令牌详情
// @Summary 获取 FRP 令牌详情
// @Tags FRP令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/token/{id} [get]
func GetFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.get")
	id := getId(c)
	frp.GetFrpTokenService(c, row, id)
}

// UpdateFrpTokenHandler 更新 FRP 令牌
// @Summary 更新 FRP 令牌
// @Tags FRP令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Param request body frp.UpdateFrpToken true "令牌参数"
// @Success 200 {object} response.Response
// @Router /frp/token/{id} [put]
func UpdateFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.update")
	id := getId(c)
	req := frpModel.UpdateFrpToken{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	frp.UpdateFrpTokenService(c, req, row, id)
}

// DeleteFrpTokenHandler 删除 FRP 令牌
// @Summary 删除 FRP 令牌
// @Tags FRP令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/token/{id} [delete]
func DeleteFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.delete")
	id := getId(c)
	frp.DeleteFrpTokenService(c, row, id)
}

// GenerateFrpTokenHandler 重新生成 FRP 令牌
// @Summary 重新生成 FRP 令牌
// @Tags FRP令牌
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Success 200 {object} response.Response
// @Router /frp/token/{id} [post]
func GenerateFrpTokenHandler(c *gin.Context) {
	row := checkPermission(c, "frp.token.generate")
	id := getId(c)
	frp.GenerateFrpTokenService(c, row, id)
}
