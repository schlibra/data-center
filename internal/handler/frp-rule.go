package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

func portCheck(c *gin.Context, min int, max int) {
	if min < 1 || max < 1 || min > 65535 || max > 65535 || min > max {
		sendI18n(c, 400, "frp.rule.port_not_valid", nil)
	}
}

// CreateFrpRuleHandler 创建 FRP 规则
// @Summary 创建 FRP 规则
// @Tags FRP规则
// @Accept application/json
// @Produce application/json
// @param request body request.CreateFrpRule true "Frp规则"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/rule/ [post]
func CreateFrpRuleHandler(c *gin.Context) {
	row := parseToken(c)
	var req request.CreateFrpRule
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	portCheck(c, req.Min, req.Max)
	service.CreateFrpRuleService(c, row, req)
}

// ListFrpRuleHandler 获取 FRP 规则列表
// @Summary 获取 FRP 规则列表
// @Tags FRP规则
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.FrpRuleTable}
// @Router /frp/rule/ [get]
func ListFrpRuleHandler(c *gin.Context) {
	row := parseToken(c)
	service.ListFrpRuleService(c, row)
}

// GetFrpRuleHandler 获取 FRP 规则详情
// @Summary 获取 FRP 规则详情
// @Tags FRP规则
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Success 200 {object} response.Response{data=models.FrpRuleTable}
// @Router /frp/rule/{id} [get]
func GetFrpRuleHandler(c *gin.Context) {
	row := parseToken(c)
	var req request.UriId
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GetFrpRuleService(c, row, req.Id)
}

// UpdateFrpRuleHandler 更新 FRP 规则
// @Summary 更新 FRP 规则
// @Tags FRP规则
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Param request body request.UpdateFrpRule true "Frp规则"
// @Success 200 {object} response.Response
// @Router /frp/rule/{id} [put]
func UpdateFrpRuleHandler(c *gin.Context) {
	row := parseToken(c)
	var req request.UpdateFrpRule
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	portCheck(c, req.Min, req.Max)
	service.UpdateFrpRuleService(c, row, req, reqId.Id)
}

// DeleteFrpRuleHandler 删除 FRP 规则
// @Summary 删除 FRP 规则
// @Tags FRP规则
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "规则ID"
// @Success 200 {object} response.Response
// @Router /frp/rule/{id} [delete]
func DeleteFrpRuleHandler(c *gin.Context) {
	row := parseToken(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteFrpRuleService(c, row, reqId.Id)
}
