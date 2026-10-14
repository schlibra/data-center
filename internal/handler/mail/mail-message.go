package mail

import (
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

func ListMailMessageHandler(c *gin.Context) {
	checkPermission(c, "mail.message.get")
	domain := getDomain(c)
	record := getRecord(c)
	mail.ListMailMessageService(c, domain, record)
}
func GetMailMessageHandler(c *gin.Context) {

}
func SendMailMessageHandler(c *gin.Context) {

}
func DeleteMailMessageHandler(c *gin.Context) {

}
