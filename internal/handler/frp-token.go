package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateFrpTokenHandler 创建 FRP 令牌
// @Summary 创建 FRP 令牌
// @Tags FRP令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.FrpTokenCreate true "令牌参数"
// @Success 200 {object} response.Response
// @Router /frp/token/ [post]
func CreateFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.create")
	var req request.FrpTokenCreate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.CreateFrpTokenService(c, req, row)
}

// ListFrpTokenHandler 获取 FRP 令牌列表
// @Summary 获取 FRP 令牌列表
// @Tags FRP令牌
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /frp/token/ [get]
func ListFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.get")
	service.ListFrpTokenService(c, row)
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
	row := parseToken(c)
	checkPermission(c, "frp.token.get")
	var req request.UriId
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GetFrpTokenService(c, row, req.Id)
}

// UpdateFrpTokenHandler 更新 FRP 令牌
// @Summary 更新 FRP 令牌
// @Tags FRP令牌
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "令牌ID"
// @Param request body request.FrpTokenUpdate true "令牌参数"
// @Success 200 {object} response.Response
// @Router /frp/token/{id} [put]
func UpdateFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.update")
	var req request.FrpTokenUpdate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}

	service.UpdateFrpTokenService(c, req, row, reqId.Id)
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
	row := parseToken(c)
	checkPermission(c, "frp.token.delete")
	var req request.UriId
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.DeleteFrpTokenService(c, row, req.Id)
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
	row := parseToken(c)
	checkPermission(c, "frp.token.generate")
	var req request.UriId
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GenerateFrpTokenService(c, row, req.Id)
}
