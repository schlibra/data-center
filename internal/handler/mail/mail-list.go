package mail

import (
	mailDomain "data-center/internal/models/request/mail"
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

func ListMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	domain := getDomain(c)
	mail.ListMailListService(c, domain)
}
func GetMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	id := getRecord(c)
	domain := getDomain(c)
	mail.GetMailListService(c, domain, id)
}
func CreateMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.create")
	domain := getDomain(c)
	req := mailDomain.CreateMailList{}.Parse(c)
	mail.CreateMailListService(c, domain, req)
}
func UpdateMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.update")
	id := getRecord(c)
	domain := getDomain(c)
	req := mailDomain.UpdateMailList{}.Parse(c)
	mail.UpdateMailListService(c, domain, id, req)
}
func DeleteMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.delete")
	id := getRecord(c)
	domain := getDomain(c)
	mail.DeleteMailListService(c, domain, id)
}
