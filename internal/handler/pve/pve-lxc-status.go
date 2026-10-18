package pve

import (
	pveModel "data-center/internal/models/request/pve"
	"data-center/internal/service/pve"

	"github.com/gin-gonic/gin"
)

func GetPveLxcStatusHandler(c *gin.Context) {
	checkAdmin(c)
	node := getNode(c)
	vmId := getVmId(c)
	pve.GetPveLxcStatusService(c, node, vmId)
}
func SetPveLxcStatusHandler(c *gin.Context) {
	checkAdmin(c)
	node := getNode(c)
	vmId := getVmId(c)
	req := pveModel.SetPveLxcStatus{}.Parse(c)
	pve.SetPveLxcStatusService(c, node, vmId, req.Action)
}
