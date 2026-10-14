package mail

import (
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

// ListMailMessageHandler 获取邮件消息列表
// @Summary 获取邮件消息列表
// @Tags 邮箱-消息
// @Produce application/json
// @Accept application/json
// @Param domain path int true "域名ID"
// @Param record path string true "邮箱记录ID"
// @Success 200 {object} response.Response
// @Router /mail/domain/{domain}/{record} [get]
func ListMailMessageHandler(c *gin.Context) {
	checkPermission(c, "mail.message.get")
	domain := getDomain(c)
	record := getRecord(c)
	mail.ListMailMessageService(c, domain, record)
}

// GetMailMessageHandler 获取邮件消息详情
// @Summary 获取邮件消息详情
// @Tags 邮箱-消息
// @Produce application/json
// @Accept application/json
// @Param domain path int true "域名ID"
// @Param record path string true "邮箱记录ID"
// @Param id path int true "邮件消息ID"
// @Success 200 {object} response.Response
// @Router /mail/domain/{domain}/{record}/{id} [get]
func GetMailMessageHandler(c *gin.Context) {
	checkPermission(c, "mail.message.get")
	domain := getDomain(c)
	record := getRecord(c)
	id := getId(c)
	mail.GetMailMessageService(c, domain, record, id)
}
