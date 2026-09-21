package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

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
func UserLogoutHandler(c *gin.Context) {
	row := parseToken(c)
	service.UserLogoutService(c, row)
}
func UserInfoHandler(c *gin.Context) {
	row := parseToken(c)
	service.UserInfoService(c, row)
}
func UserUpdateHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "user.update")
	var req request.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.UserUpdateService(c, req, row)
}
func UserPasswordHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "user.password")
	var req request.UserPassword
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.UserPasswordService(c, req, row)
}
