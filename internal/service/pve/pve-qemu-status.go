package pve

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPveQemuStatusService(c *gin.Context, node string, vmId int) {
	var res struct {
		Data H `json:"data" `
	}
	req, url, err := pveRequest("/nodes/"+node+"/qemu/"+strconv.Itoa(vmId)+"/status/current", &res)
	if err != nil {
		sendError(c, err)
	}
	if _, err := req.Get(url); err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res.Data)
}
func SetPveQemuStatusService(c *gin.Context, node string, vmId int, action string) {
	var res struct {
		Data string `json:"data"`
	}
	req, url, err := pveRequest("/nodes/"+node+"/qemu/"+strconv.Itoa(vmId)+"/status/"+action, &res)
	if err != nil {
		sendError(c, err)
	}
	resp, err := req.SetBody("").Post(url)
	if err != nil {
		sendError(c, err)
	}
	fmt.Printf("%v\n", resp)
	sendJson(c, 200, "success", res.Data)
}
