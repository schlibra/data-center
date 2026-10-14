package mail

import (
	mailDomain "data-center/internal/models/request/mail"
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

// ListMailListHandler 获取邮箱列表
// @Summary 获取邮箱列表
// @Tags 邮箱-邮箱管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param domain path int true "域名ID"
// @Success 200 {object} response.Response
// @Router /mail/list/{domain} [get]
func ListMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	domain := getDomain(c)
	mail.ListMailListService(c, domain)
}

// GetMailListHandler 获取邮箱详情
// @Summary 获取邮箱详情
// @Tags 邮箱-邮箱管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param domain path int true "域名ID"
// @Param id path int true "邮箱ID"
// @Success 200 {object} response.Response
// @Router /mail/list/{domain}/{id} [get]
func GetMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	id := getRecord(c)
	domain := getDomain(c)
	mail.GetMailListService(c, domain, id)
}

// CreateMailListHandler 创建邮箱
// @Summary 创建邮箱
// @Tags 邮箱-邮箱管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param domain path int true "域名ID"
// @Param request body mailDomain.CreateMailList true "邮箱创建请求"
// @Success 200 {object} response.Response
// @Router /mail/list/{domain} [post]
func CreateMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.create")
	domain := getDomain(c)
	req := mailDomain.CreateMailList{}.Parse(c)
	mail.CreateMailListService(c, domain, req)
}

// UpdateMailListHandler 更新邮箱
// @Summary 更新邮箱
// @Tags 邮箱-邮箱管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param domain path int true "域名ID"
// @Param record path string true "邮箱记录ID"
// @Param request body mailDomain.UpdateMailList true "邮箱更新请求"
// @Success 200 {object} response.Response
// @Router /mail/list/{domain}/{id} [put]
func UpdateMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.update")
	domain := getDomain(c)
	record := getRecord(c)
	req := mailDomain.UpdateMailList{}.Parse(c)
	mail.UpdateMailListService(c, domain, record, req)
}

// DeleteMailListHandler 删除邮箱
// @Summary 删除邮箱
// @Tags 邮箱-邮箱管理
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param domain path int true "域名ID"
// @Param record path string true "邮箱记录ID"
// @Success 200 {object} response.Response
// @Router /mail/list/{domain}/{id} [delete]
func DeleteMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.delete")
	domain := getDomain(c)
	record := getRecord(c)
	mail.DeleteMailListService(c, domain, record)
}
