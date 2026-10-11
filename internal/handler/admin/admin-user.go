package admin

import (
	adminModel "data-center/internal/models/request/admin"
	"data-center/internal/service/admin"

	"github.com/gin-gonic/gin"
)

// ListAdminUserHandler 获取用户列表（管理员）
// @summary 获取用户列表（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @success 200 {object} response.Response{data=[]models.UserTable}
// @router /admin/user [get]
func ListAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	admin.ListAdminUserService(c)
}

// GetAdminUserHandler 获取用户详情（管理员）
// @summary 获取用户详情（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户ID"
// @success 200 {object} response.Response{data=models.UserTable}
// @router /admin/user/{id} [get]
func GetAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	admin.GetAdminUserService(c, id)
}

// CreateAdminUserHandler 创建用户（管理员）
// @summary 创建用户（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body admin.CreateAdminUser true "用户信息"
// @success 200 {object} response.Response
// @router /admin/user [post]
func CreateAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req adminModel.CreateAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	if req.Username == "" {
		sendI18n(c, 400, "admin.user.username_empty", nil)
	}
	if req.Password == "" {
		sendI18n(c, 400, "admin.user.password_empty", nil)
	}
	id := getId(c)
	admin.CreateAdminUserService(c, req, id)
}

// UpdateAdminUserHandler 更新用户（管理员）
// @summary 更新用户（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body admin.UpdateAdminUser true "用户信息"
// @param id path int true "用户ID"
// @success 200 {object} response.Response
// @router /admin/user/{id} [put]
func UpdateAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req adminModel.UpdateAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	id := getId(c)
	admin.UpdateAdminUserService(c, req, id)
}

// PasswordAdminUserHandler 修改用户密码（管理员）
// @summary 修改用户密码（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body admin.PasswordAdminUser true "用户密码信息"
// @param id path int true "用户ID"
// @success 200 {object} response.Response
// @router /admin/user/{id} [patch]
func PasswordAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req adminModel.PasswordAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	if req.Password == "" {
		sendI18n(c, 400, "admin.user.password_empty", nil)
	}
	id := getId(c)
	admin.PasswordAdminUserService(c, req, id)
}

// DeleteAdminUserHandler 删除用户（管理员）
// @summary 删除用户（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param id path int true "用户ID"
// @success 200 {object} response.Response
// @router /admin/user/{id} [delete]
func DeleteAdminUserHandler(c *gin.Context) {
	row := checkAdmin(c)
	id := getId(c)
	admin.DeleteAdminUserService(c, row, id)
}
