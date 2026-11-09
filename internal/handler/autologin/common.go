package autologin

import (
	"data-center/internal/handler"
	"data-center/internal/models"

	"github.com/gin-gonic/gin"
)

type H map[string]any

func checkPermission(c *gin.Context, pmsKey string) models.UserTable {
	return handler.CheckPermission(c, pmsKey)
}
func checkAdmin(c *gin.Context) models.UserTable {
	return handler.CheckAdmin(c)
}
func getId(c *gin.Context) int {
	return handler.GetId(c)
}
