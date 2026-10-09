package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

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
	service.ListAdminSettingsService(c)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.GetAdminSettingsService(c, reqId.Id)
}

// CreateAdminSettingsHandler 创建设置
// @Summary 创建设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.CreateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings [post]
func CreateAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	if req.Key == "" {
		sendI18n(c, 400, "settings.key_empty", nil)
	}
	service.CreateAdminSettingsService(c, req)
}

// UpdateAdminSettingsHandler 更新设置
// @Summary 更新设置
// @Summary 更新设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param id path int true "设置ID"
// @Param request body request.UpdateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings/{id} [put]
func UpdateAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req request.UpdateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.UpdateAdminSettingsService(c, reqId.Id, req)
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
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	service.DeleteAdminSettingsService(c, reqId.Id)
}

// SetAdminSettingsHandler 写入设置
// @Summary 写入设置
// @Tags 管理员-设置
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param request body request.CreateAdminSettings true "设置数据"
// @Success 200 {object} response.Response
// @Router /admin/settings [put]
func SetAdminSettingsHandler(c *gin.Context) {
	checkAdmin(c)
	var req request.CreateAdminSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	service.SetAdminSettingsService(c, req)
}
