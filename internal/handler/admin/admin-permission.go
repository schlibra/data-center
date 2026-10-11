package admin

import (
	adminModel "data-center/internal/models/request/admin"
	"data-center/internal/service/admin"

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
	admin.ListAdminPermissionService(c)
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
	id := getId(c)
	admin.GetAdminPermissionService(c, id)
}

// CreateAdminPermissionHandler 创建权限（管理员）
// @summary 创建权限（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body admin.CreateAdminPermission true "权限信息"
// @success 200 {object} response.Response
// @router /admin/permission [post]
func CreateAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	req := adminModel.CreateAdminPermission{}.Parse(c)
	if req.Key == "" || req.Name == "" {
		sendI18n(c, 400, "admin.permission.key_name_empty", nil)
	}
	admin.CreateAdminPermissionService(c, req)
}

// UpdateAdminPermissionHandler 更新权限（管理员）
// @summary 更新权限（管理员）
// @tags 管理员-权限管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "权限ID"
// @param request body admin.UpdateAdminPermission true "权限信息"
// @success 200 {object} response.Response
// @router /admin/permission/{id} [put]
func UpdateAdminPermissionHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	req := adminModel.UpdateAdminPermission{}.Parse(c)
	if req.Key == "" || req.Name == "" {
		sendI18n(c, 400, "admin.permission.key_name_empty", nil)
	}
	if req.Parent == id {
		sendI18n(c, 400, "admin.permission.id_equal_parent", nil)
	}
	admin.UpdateAdminPermissionService(c, id, req)
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
	id := getId(c)
	admin.DeleteAdminPermissionService(c, id)
}
