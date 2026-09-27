package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// ListOpenVPNGroupHandler 获取OpenVPN IP组列表
// @Summary 获取OpenVPN IP组列表
// @Tags OpenVPN-IP组管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=response.OpenVPNGroupResult}
// @Router /openvpn/group [get]
func ListOpenVPNGroupHandler(c *gin.Context) {
	checkPermission(c, "openvpn.group.get")
	service.ListOpenVPNGroupService(c)
}

// GetOpenVPNGroupHandler 获取OpenVPN IP组详情
// @Summary 获取OpenVPN IP组详情
// @Tags OpenVPN-IP组管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "IP组ID"
// @Success 200 {object} response.Response{data=response.OpenVPNGroupResult}
// @Router /openvpn/group/{id} [get]
func GetOpenVPNGroupHandler(c *gin.Context) {
	checkPermission(c, "openvpn.group.get")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetOpenVPNGroupService(c, reqId.Id)
}

// CreateOpenVPNGroupHandler 创建OpenVPN IP组
// @Summary 创建OpenVPN IP组
// @Tags OpenVPN-IP组管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param request body request.CreateOpenVPNGroup true "IP组参数"
// @Success 200 {object} response.Response
// @Router /openvpn/group [post]
func CreateOpenVPNGroupHandler(c *gin.Context) {
	checkPermission(c, "openvpn.group.create")
	var req request.CreateOpenVPNGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.CreateOpenVPNGroupService(c, req)
}

// UpdateOpenVPNGroupHandler 更新OpenVPN IP组
// @Summary 更新OpenVPN IP组
// @Tags OpenVPN-IP组管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "IP组ID"
// @Param request body request.UpdateOpenVPNGroup true "IP组参数"
// @Success 200 {object} response.Response
// @Router /openvpn/group/{id} [put]
func UpdateOpenVPNGroupHandler(c *gin.Context) {
	checkPermission(c, "openvpn.group.update")
	var req request.UpdateOpenVPNGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.UpdateOpenVPNGroupService(c, reqId.Id, req)
}

// DeleteOpenVPNGroupHandler 删除OpenVPN IP组
// @Summary 删除OpenVPN IP组
// @Tags OpenVPN-IP组管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "IP组ID"
// @Success 200 {object} response.Response
// @Router /openvpn/group/{id} [delete]
func DeleteOpenVPNGroupHandler(c *gin.Context) {
	checkPermission(c, "openvpn.group.delete")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteOpenVPNGroupService(c, reqId.Id)
}
