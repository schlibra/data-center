package autologin

import (
	autologinModel "data-center/internal/models/request/autologin"
	"data-center/internal/service/autologin"

	"github.com/gin-gonic/gin"
)

func ListAdminAutoLoginHandler(c *gin.Context) {
	checkAdmin(c)
	autologin.ListAdminAutoLoginService(c)
}
func GetAdminAutoLoginHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	autologin.GetAdminAutoLoginService(c, id)
}
func CreateAdminAutoLoginHandler(c *gin.Context) {
	checkAdmin(c)
	caal := autologinModel.CreateAdminAutoLogin{}.Parse(c)
	autologin.CreateAdminAutoLoginService(c, caal)
}
func UpdateAdminAutoLoginHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	uaal := autologinModel.UpdateAdminAutoLogin{}.Parse(c)
	autologin.UpdateAdminAutoLoginService(c, id, uaal)
}
func DeleteAdminAutoLoginHandler(c *gin.Context) {
	checkAdmin(c)
	id := getId(c)
	autologin.DeleteAdminAutoLoginService(c, id)
}
