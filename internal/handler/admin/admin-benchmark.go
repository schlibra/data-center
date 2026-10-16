package admin

import (
	"data-center/internal/service/admin"

	"github.com/gin-gonic/gin"
)

func DoAdminBenchmarkHandler(c *gin.Context) {
	checkAdmin(c)
	admin.DoAdminBenchmarkService(c)
}
