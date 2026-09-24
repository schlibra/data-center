package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// ListAdminPermissionsHandler 获取权限列表（管理员）
// @summary 获取权限列表（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @success 200 {object} response.Response{data=[]models.PermissionTable}
// @router /admin/permission [get]
func ListAdminPermissionsHandler(c *gin.Context) {
	checkAdmin(c)
	service.ListAdminPermissionService(c)
}

// GetAdminPermissionHandler 获取权限详情（管理员）
// @summary 获取权限详情（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "权限ID"
// @success 200 {object} response.Response{data=models.PermissionTable}
// @router /admin/permission/{id} [get]
func GetAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetAdminPermissionService(c, reqId.Id)
}

// CreateAdminPermissionHandler 创建权限（管理员）
// @summary 创建权限（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body request.CreateAdminPermission true "权限信息"
// @success 200 {object} response.Response
// @router /admin/permission [post]
func CreateAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateAdminPermission
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Key == "" || req.Name == "" {
		sendI18n(c, 400, "admin.permission.key_name_empty", nil)
	}
	service.CreateAdminPermissionService(c, req)
}

// UpdateAdminPermissionHandler 更新权限（管理员）
// @summary 更新权限（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "权限ID"
// @param request body request.UpdateAdminPermission true "权限信息"
// @success 200 {object} response.Response
// @router /admin/permission/{id} [put]
func UpdateAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req request.UpdateAdminPermission
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Key == "" || req.Name == "" {
		sendI18n(c, 400, "admin.permission.key_name_empty", nil)
	}
	service.UpdateAdminPermissionService(c, reqId.Id, req)
}

// DeleteAdminPermissionHandler 删除权限（管理员）
// @summary 删除权限（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "权限ID"
// @success 200 {object} response.Response
// @router /admin/permission/{id} [delete]
func DeleteAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteAdminPermissionService(c, reqId.Id)
}
