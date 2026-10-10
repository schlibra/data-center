package openvpn

import (
	"data-center/internal/models/request"
	openvpnModel "data-center/internal/models/request/openvpn"
	"data-center/internal/service/openvpn"

	"github.com/gin-gonic/gin"
)

// ListOpenVPNUserHandler 获取OpenVPN用户列表
// @Summary 获取OpenVPN用户列表
// @Tags OpenVPN-用户管理
// @Produce application/json
// @Accpept application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=response.OpenVPNUser}
// @Router /openvpn/user [GET]
func ListOpenVPNUserHandler(c *gin.Context) {
	checkPermission(c, "openvpn.user.get")
	openvpn.ListOpenVPNUserService(c)
}

// GetOpenVPNUserHandler 获取OpenVPN用户详情
// @Summary 获取OpenVPN用户详情
// @Tags OpenVPN-用户管理
// @Produce application/json
// @Accpept application/json
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=response.OpenVPNUser}
// @Router /openvpn/user/{id} [GET]
func GetOpenVPNUserHandler(c *gin.Context) {
	checkPermission(c, "openvpn.user.get")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	openvpn.GetOpenVPNUserService(c, reqId.Id)
}

// CreateOpenVPNUserHandler 创建OpenVPN用户
// @Summary 创建OpenVPN用户
// @Tags OpenVPN-用户管理
// @Produce application/json
// @Accpept application/json
// @Security BearerAuth
// @Param request body openvpn.CreateOpenVPNUser true "用户信息"
// @Success 200 {object} response.Response{data=response.OpenVPNUserChange}
// @Router /openvpn/user [POST]
func CreateOpenVPNUserHandler(c *gin.Context) {
	checkPermission(c, "openvpn.user.create")
	var req openvpnModel.CreateOpenVPNUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	openvpn.CreateOpenVPNUserService(c, req)
}

// UpdateOpenVPNUserHandler 更新OpenVPN用户
// @Summary 更新OpenVPN用户
// @Tags OpenVPN-用户管理
// @Produce application/json
// @Accpept application/json
// @Security BearerAuth
// @Param request body openvpn.CreateOpenVPNUser true "用户信息"
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=response.OpenVPNUserChange}
// @Router /openvpn/user/{id} [PUT]
func UpdateOpenVPNUserHandler(c *gin.Context) {
	checkPermission(c, "openvpn.user.create")
	var req openvpnModel.UpdateOpenVPNUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	openvpn.UpdateOpenVPNUserService(c, reqId.Id, req)
}

// DeleteOpenVPNUserHandler 删除OpenVPN用户
// @Summary 删除OpenVPN用户
// @Tags OpenVPN-用户管理
// @Produce application/json
// @Accpept application/json
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /openvpn/user/{id} [DELETE]
func DeleteOpenVPNUserHandler(c *gin.Context) {
	checkPermission(c, "openvpn.user.delete")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	openvpn.DeleteOpenVPNUserService(c, reqId.Id)
}
