package handler

import (
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// GetAdminVersionHandler 获取程序版本信息
// @Summary 获取程序版本信息
// @Tags 管理员-版本
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /admin/version [get]
func GetAdminVersionHandler(c *gin.Context, version string, commit, buildTime string) {
	checkAdmin(c)
	service.GetAdminVersionService(c, version, commit, buildTime)
}

// UpgradeAdminVersionHandler 升级程序版本
// @Summary 升级程序版本
// @Tags 管理员-版本
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /admin/version [post]
func UpgradeAdminVersionHandler(c *gin.Context) {
	checkAdmin(c)
	service.UpgradeAdminVersionService(c)
}
