package pve

import (
	pveModel "data-center/internal/models/request/pve"
	"data-center/internal/service/pve"

	"github.com/gin-gonic/gin"
)

func GetPveQemuStatusHandler(c *gin.Context) {
	checkAdmin(c)
	node := getNode(c)
	vmId := getVmId(c)
	pve.GetPveQemuStatusService(c, node, vmId)
}

func SetPveQemuStatusHandler(c *gin.Context) {
	checkAdmin(c)
	node := getNode(c)
	vmId := getVmId(c)
	req := pveModel.SetPveQemuStatus{}.Parse(c)
	pve.SetPveQemuStatusService(c, node, vmId, req.Action)
}
