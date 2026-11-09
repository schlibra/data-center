package autologin

import (
	"data-center/internal/service/autologin"

	"github.com/gin-gonic/gin"
)

func ListAutoLoginHandler(c *gin.Context) {
	checkPermission(c, "autologin.get")
	autologin.ListAutoLoginService(c)
}
func GetAutoLoginHandler(c *gin.Context) {
	checkPermission(c, "autologin.get")
	id := getId(c)
	autologin.GetAutoLoginService(c, id)
}
