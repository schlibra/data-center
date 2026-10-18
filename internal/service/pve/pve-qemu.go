package pve

import (
	"data-center/internal/models/response"

	"github.com/gin-gonic/gin"
)

func ListPveQemuService(c *gin.Context, node string) {
	var res response.PveQemu
	req, url, err := pveRequest("/nodes/"+node+"/qemu", &res)
	if err != nil {
		sendError(c, err)
	}
	_, err = req.Get(url)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res.Data)
}
