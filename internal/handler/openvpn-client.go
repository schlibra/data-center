package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// ListOpenVPNClientHandler 获取OpenVPN客户端列表
// @Summary 获取OpenVPN客户端列表
// @Tags OpenVPN-客户端管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /openvpn/client [get]
func ListOpenVPNClientHandler(c *gin.Context) {
	checkPermission(c, "openvpn.client.get")
	service.ListOpenVPNClientService(c)
}

// GetOpenVPNClientHandler 获取OpenVPN客户端详情
// @Summary 获取OpenVPN客户端详情
// @Tags OpenVPN-客户端管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "客户端ID"
// @Success 200 {object} response.Response
// @Router /openvpn/client/{id} [get]
func GetOpenVPNClientHandler(c *gin.Context) {
	checkPermission(c, "openvpn.client.get")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetOpenVPNClientService(c, reqId.Id)
}

// KickOpenVPNClientHandler 踢出OpenVPN客户端
// @Summary 踢出OpenVPN客户端
// @Tags OpenVPN-客户端管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "客户端ID"
// @Success 200 {object} response.Response
// @Router /openvpn/client/{id} [delete]
func KickOpenVPNClientHandler(c *gin.Context) {
	checkPermission(c, "openvpn.client.delete")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.KickOpenVPNClientService(c, reqId.Id)
}
