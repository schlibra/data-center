package pve

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, err error) {
	sendError(c, err)
}

func checkAdmin(c *gin.Context) {
	handler.CheckAdmin(c)
}

func getNode(c *gin.Context) string {
	type ReqNode struct {
		Node string `uri:"node"`
	}
	var reqNode ReqNode
	if err := c.ShouldBindUri(&reqNode); err != nil {
		sendError(c, err)
	}
	return reqNode.Node
}
func getVmId(c *gin.Context) int {
	type ReqVmId struct {
		VmId int `uri:"vmid"`
	}
	var reqVmId ReqVmId
	if err := c.ShouldBindUri(&reqVmId); err != nil {
		sendError(c, err)
	}
	return reqVmId.VmId
}
