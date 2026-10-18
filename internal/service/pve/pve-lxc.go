package pve

import "github.com/gin-gonic/gin"

func ListPveLxcService(c *gin.Context, node string) {
	var res struct {
		Data []H `json:"data"`
	}
	req, url, err := pveRequest("/nodes/"+node+"/lxc", &res)
	if err != nil {
		sendError(c, err)
	}
	_, err = req.Get(url)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res.Data)
}
