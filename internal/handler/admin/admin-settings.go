package admin

import (
	adminModel "data-center/internal/models/request/admin"
	"data-center/internal/service/admin"

	"github.com/gin-gonic/gin"
)

// ListAdminSettingsHandler 获取所有设置
// @Summary 获取所有设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SettingsTable}
// @Router /admin/settings [get]
func ListAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	admin.ListAdminSettingsService(c)
}

// GetAdminSettingsHandler 获取单个设置
// @Summary 获取单个设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Param id path int true "设置ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=models.SettingsTable}
// @Router /admin/settings/{id} [get]
func GetAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	admin.GetAdminSettingsService(c, id)
}

// CreateAdminSettingsHandler 创建设置
// @Summary 创建设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body admin.CreateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings [post]
func CreateAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	var req adminModel.CreateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Key == "" {
		sendI18n(c, 400, "settings.key_empty", nil)
	}
	admin.CreateAdminSettingsService(c, req)
}

// UpdateAdminSettingsHandler 更新设置
// @Summary 更新设置
// @Summary 更新设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "设置ID"
// @Param request body admin.UpdateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings/{id} [put]
func UpdateAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	var req adminModel.UpdateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	admin.UpdateAdminSettingsService(c, id, req)
}

// DeleteAdminSettingsHandler 删除设置
// @Summary 删除设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "设置ID"
// @Success 200 {object} response.Response
// @Router /admin/settings/{id} [delete]
func DeleteAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	admin.DeleteAdminSettingsService(c, id)
}

// SetAdminSettingsHandler 写入设置
// @Summary 写入设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body admin.CreateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings [put]
func SetAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	var req adminModel.CreateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	admin.SetAdminSettingsService(c, req)
}
