package mail

import (
	mailModel "data-center/internal/models/request/mail"
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

// ListMailDomainHandler 获取邮箱域名列表
// @Summary 获取邮箱域名列表
// @Tags 邮箱-域名
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.MailDomainTable}
// @Router /mail/domain [get]
func ListMailDomainHandler(c *gin.Context) {
	checkPermission(c, "mail.domain.get")
	mail.ListMailDomainService(c)
}

// GetMailDomainHandler 获取邮箱域名详情
// @Summary 获取邮箱域名详情
// @Tags 邮箱-域名
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "邮箱域名ID"
// @Success 200 {object} response.Response{data=models.MailDomainTable}
// @Router /mail/domain/{id} [get]
func GetMailDomainHandler(c *gin.Context) {
	checkPermission(c, "mail.domain.get")
	id := getId(c)
	mail.GetMailDomainService(c, id)
}

// CreateMailDomainHandler 创建邮箱域名
// @Summary 创建邮箱域名
// @Tags 邮箱-域名
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param request body mailModel.CreateMailDomain true "邮箱域名表"
// @Success 200 {object} response.Response
// @Router /mail/domain [post]
func CreateMailDomainHandler(c *gin.Context) {
	checkPermission(c, "mail.domain.create")
	req := mailModel.CreateMailDomain{}.Parse(c)
	mail.CreateMailDomainService(c, req)
}

// UpdateMailDomainHandler 更新邮箱域名
// @Summary 更新邮箱域名
// @Tags 邮箱-域名
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "邮箱域名ID"
// @Param request body mailModel.UpdateMailDomain true "邮箱域名表"
// @Success 200 {object} response.Response
// @Router /mail/domain/{id} [put]
func UpdateMailDomainHandler(c *gin.Context) {
	checkPermission(c, "mail.domain.update")
	id := getId(c)
	req := mailModel.UpdateMailDomain{}.Parse(c)
	mail.UpdateMailDomainService(c, id, req)
}

// DeleteMailDomainHandler 删除邮箱域名
// @Summary 删除邮箱域名
// @Tags 邮箱-域名
// @Produce application/json
// @Accept application/json
// @Security BearerAuth
// @Param id path int true "邮箱域名ID"
// @Success 200 {object} response.Response
// @Router /mail/domain/{id} [delete]
func DeleteMailDomainHandler(c *gin.Context) {
	checkPermission(c, "mail.domain.delete")
	id := getId(c)
	mail.DeleteMailDomainService(c, id)
}
