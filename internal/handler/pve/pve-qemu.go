package pve

import (
	"data-center/internal/service/pve"

	"github.com/gin-gonic/gin"
)

func ListPveQemuHandler(c *gin.Context) {
	checkAdmin(c)
	node := getNode(c)
	pve.ListPveQemuService(c, node)
}
