package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

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
	service.ListAdminUserService(c)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
		return
	}
	service.GetAdminUserService(c, reqId.Id)
}

// CreateAdminUserHandler 创建用户（管理员）
// @summary 创建用户（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body request.CreateAdminUser true "用户信息"
// @success 200 {object} response.Response
// @router /admin/user [post]
func CreateAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
		return
	}
	service.CreateAdminUserService(c, req, reqId.Id)
}

// UpdateAdminUserHandler 更新用户（管理员）
// @summary 更新用户（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body request.UpdateAdminUser true "用户信息"
// @param id path int true "用户ID"
// @success 200 {object} response.Response
// @router /admin/user/{id} [put]
func UpdateAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.UpdateAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
		return
	}
	service.UpdateAdminUserService(c, req, reqId.Id)
}

// PasswordAdminUserHandler 修改用户密码（管理员）
// @summary 修改用户密码（管理员）
// @tags 管理员-用户管理
// @accept application/json
// @produce application/json
// @security BearerAuth
// @param request body request.PasswordAdminUser true "用户密码信息"
// @param id path int true "用户ID"
// @success 200 {object} response.Response
// @router /admin/user/{id}/password [patch]
func PasswordAdminUserHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.PasswordAdminUser
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
		return
	}
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
		return
	}
	service.PasswordAdminUserService(c, req, reqId.Id)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
		return
	}
	service.DeleteAdminUserService(c, row, reqId.Id)
}
