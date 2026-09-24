package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

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
	service.ListFrpAdminTokenService(c)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetFrpAdminTokenService(c, reqId.Id)
}

// CreateFrpAdminTokenHandler 创建 FRP 令牌（管理员）
// @Summary 创建 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.CreateFrpAdminToken true "令牌信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/ [post]
func CreateFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateFrpAdminToken
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	service.CreateFrpAdminTokenService(c, req)
}

// UpdateFrpAdminTokenHandler 更新 FRP 令牌（管理员）
// @Summary 更新 FRP 令牌（管理员）
// @Tags FRP管理-令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Param request body request.UpdateFrpAdminToken true "令牌信息"
// @Success 200 {object} response.Response
// @Router /frp/admin/token/{id} [put]
func UpdateFrpAdminTokenHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req request.UpdateFrpAdminToken
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Name == "" {
		sendI18n(c, 400, "frp.token.name_empty", nil)
	}
	service.UpdateFrpAdminTokenService(c, reqId.Id, req)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteFrpAdminTokenService(c, reqId.Id)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GenerateFrpAdminTokenService(c, reqId.Id)
}
