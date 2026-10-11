package mail

import (
	"data-center/internal/service/mail"

	"github.com/gin-gonic/gin"
)

func ListMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	mail.ListMailListService(c)
}
func GetMailListHandler(c *gin.Context) {
	checkPermission(c, "mail.list.get")
	id := getId(c)
	mail.GetMailListService(c, id)
}
func CreateMailListHandler(c *gin.Context) {

}
func UpdateMailListHandler(c *gin.Context) {

}
func DeleteMailListHandler(c *gin.Context) {

}
