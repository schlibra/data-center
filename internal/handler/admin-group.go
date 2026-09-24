package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// ListAdminGroupHandler 获取用户组列表（管理员）
// @Summary 获取用户组列表（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @success 200 {object} response.Response{data=[]models.GroupTable}
// @router /admin/group [get]
func ListAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	service.ListAdminGroupService(c)
}

// GetAdminGroupHandler 获取用户组详情（管理员）
// @Summary 获取用户组详情（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户组ID"
// @success 200 {object} response.Response{data=models.GroupTable}
// @router /admin/group/{id} [get]
func GetAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetAdminGroupService(c, reqId.Id)
}

// CreateAdminGroupHandler 创建用户组（管理员）
// @Summary 创建用户组（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body request.CreateAdminGroup true "用户组信息"
// @success 200 {object} response.Response
// @router /admin/group [post]
func CreateAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateAdminGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.CreateAdminGroupService(c, req)
}

// UpdateAdminGroupHandler 更新用户组（管理员）
// @Summary 更新用户组（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户组ID"
// @param request body request.UpdateAdminGroup true "用户组信息"
// @success 200 {object} response.Response
// @router /admin/group/{id} [put]
func UpdateAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req request.UpdateAdminGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.UpdateAdminGroupService(c, reqId.Id, req)
}

// DeleteAdminGroupHandler 删除用户组（管理员）
// @Summary 删除用户组（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户组ID"
// @success 200 {object} response.Response
// @router /admin/group/{id} [delete]
func DeleteAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteAdminGroupService(c, reqId.Id)
}
