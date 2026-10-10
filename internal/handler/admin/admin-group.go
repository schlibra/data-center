package admin

import (
	"data-center/internal/models/request"
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	admin.GetAdminGroupService(c, reqId.Id)
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
	var req adminModel.CreateAdminGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req adminModel.UpdateAdminGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Name == "" {
		sendI18n(c, 400, "admin.group.name_empty", nil)
	}
	admin.UpdateAdminGroupService(c, reqId.Id, req)
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
	admin.DeleteAdminGroupService(c, reqId.Id)
}
