package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// UserLoginHandler 用户登录
// @Summary 用户登录
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param request body request.UserLogin true "登录参数"
// @Success 200 {object} response.Response
// @Router /user/login [post]
func UserLoginHandler(c *gin.Context) {
	var req request.UserLogin
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	if req.Username == "" || req.Password == "" {
		sendI18n(c, 400, "user.login.empty", nil)
	}
	service.UserLoginService(c, req)
}

// UserLoginKeyHandler 获取登录公钥
// @Summary 获取登录公钥
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param request body request.UserLoginKey true "用户名"
// @Success 200 {object} response.Response
// @Router /user/login [put]
func UserLoginKeyHandler(c *gin.Context) {
	var req request.UserLoginKey
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	if req.Username == "" {
		sendI18n(c, 400, "user.login.username_empty", nil)
	}
	service.UserLoginKeyService(c, req)
}

// UserRegisterHandler 用户注册
// @Summary 用户注册
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param request body request.UserRegister true "注册参数"
// @Success 200 {object} response.Response
// @Router /user/register [post]
func UserRegisterHandler(c *gin.Context) {
	var req request.UserRegister
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	if req.Username == "" || req.Password == "" || req.Nickname == "" {
		sendI18n(c, 400, "user.register.empty", nil)
	}
	service.UserRegisterService(c, req)
}

// UserRegisterKeyHandler 获取注册公钥
// @Summary 获取注册公钥
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param request body request.UserRegisterKey true "用户名"
// @Success 200 {object} response.Response
// @Router /user/register [put]
func UserRegisterKeyHandler(c *gin.Context) {
	var req request.UserRegisterKey
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	if req.Username == "" {
		sendI18n(c, 400, "user.register.username_empty", nil)
	}
	service.UserRegisterKeyService(c, req)
}

// UserLogoutHandler 用户登出
// @Summary 用户登出
// @Tags 用户
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /user/logout [post]
func UserLogoutHandler(c *gin.Context) {
	row := parseToken(c)
	service.UserLogoutService(c, row)
}

// UserInfoHandler 获取当前用户信息
// @Summary 获取当前用户信息
// @Tags 用户
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=models.UserTable}
// @Router /user/ [get]
func UserInfoHandler(c *gin.Context) {
	row := parseToken(c)
	service.UserInfoService(c, row)
}

// UserUpdateHandler 更新用户信息
// @Summary 更新用户信息
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.UserUpdate true "更新参数"
// @Success 200 {object} response.Response
// @Router /user/ [put]
func UserUpdateHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "user.update")
	var req request.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.UserUpdateService(c, req, row)
}

// UserPasswordHandler 修改用户密码
// @Summary 修改用户密码
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.UserPassword true "新密码"
// @Success 200 {object} response.Response
// @Router /user/ [patch]
func UserPasswordHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "user.password")
	var req request.UserPassword
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.UserPasswordService(c, req, row)
}
