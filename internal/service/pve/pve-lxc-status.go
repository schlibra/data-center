package pve

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPveLxcStatusService(c *gin.Context, node string, vmId int) {
	var res struct {
		Data H `json:"data"`
	}
	req, url, err := pveRequest("/nodes/"+node+"/lxc/"+strconv.Itoa(vmId)+"/status/current", &res)
	if err != nil {
		sendError(c, err)
	}
	if _, err := req.Get(url); err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res.Data)
}
func SetPveLxcStatusService(c *gin.Context, node string, vmId int, action string) {
	var res H
	req, url, err := pveRequest("/nodes/"+node+"/lxc/"+strconv.Itoa(vmId)+"/status/"+action, &res)
	if err != nil {
		sendError(c, err)
	}
	if _, err := req.SetBody("").Post(url); err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res)
}
