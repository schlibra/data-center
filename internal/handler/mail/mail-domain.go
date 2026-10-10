package mail

import (
	"data-center/internal/models/request"
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
	checkPermission(c, "mail.get")
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
	checkPermission(c, "mail.get")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	mail.GetMailDomainService(c, reqId.Id)
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
	checkPermission(c, "mail.create")
	var req mailModel.CreateMailDomain
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
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
	checkPermission(c, "mail.update")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	var req mailModel.UpdateMailDomain
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, err)
	}
	mail.UpdateMailDomainService(c, reqId.Id, req)
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
	checkPermission(c, "mail.delete")
	var reqId request.UriId
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	mail.DeleteMailDomainService(c, reqId.Id)
}
