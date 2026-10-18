package pve

import (
	"data-center/internal/service/pve"

	"github.com/gin-gonic/gin"
)

func ListPveNodeHandler(c *gin.Context) {
	checkAdmin(c)
	pve.ListPveNodeService(c)
}
