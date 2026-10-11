package admin

import (
	adminModel "data-center/internal/models/request/admin"
	"data-center/internal/service/admin"

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
	admin.ListAdminGroupService(c)
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
	id := getId(c)
	admin.GetAdminGroupService(c, id)
}

// CreateAdminGroupHandler 创建用户组（管理员）
// @Summary 创建用户组（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body admin.CreateAdminGroup true "用户组信息"
// @success 200 {object} response.Response
// @router /admin/group [post]
func CreateAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	req := adminModel.CreateAdminGroup{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "admin.group.name_empty", nil)
	}
	admin.CreateAdminGroupService(c, req)
}

// UpdateAdminGroupHandler 更新用户组（管理员）
// @Summary 更新用户组（管理员）
// @tags 管理员-用户组管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户组ID"
// @param request body admin.UpdateAdminGroup true "用户组信息"
// @success 200 {object} response.Response
// @router /admin/group/{id} [put]
func UpdateAdminGroupHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	req := adminModel.UpdateAdminGroup{}.Parse(c)
	if req.Name == "" {
		sendI18n(c, 400, "admin.group.name_empty", nil)
	}
	admin.UpdateAdminGroupService(c, id, req)
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
	id := getId(c)
	admin.DeleteAdminGroupService(c, id)
}
