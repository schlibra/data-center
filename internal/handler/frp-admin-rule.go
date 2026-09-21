package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

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
	service.ListFrpAdminRuleService(c)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetFrpAdminRuleService(c, reqId.Id)
}

// CreateFrpAdminRuleHandler 创建 FRP 规则（管理员）
// @Summary 创建 FRP 规则（管理员）
// @Tags FRP管理-规则
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/admin/rule/ [post]
func CreateFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateFrpAdminRule
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	portCheck(c, req.Min, req.Max)
	service.CreateFrpAdminRuleService(c, req)
}

// UpdateFrpAdminRuleHandler 更新 FRP 规则（管理员）
// @Summary 更新 FRP 规则（管理员）
// @Tags FRP管理-规则
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Success 200 {object} response.Response
// @Router /frp/admin/rule/{id} [put]
func UpdateFrpAdminRuleHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.UpdateFrpAdminRule
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	portCheck(c, req.Min, req.Max)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.UpdateFrpAdminRuleService(c, req, reqId.Id)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteFrpAdminRuleService(c, reqId.Id)
}
