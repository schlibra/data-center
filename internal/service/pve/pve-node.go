package pve

import (
	"data-center/internal/models/response"

	"github.com/gin-gonic/gin"
)

func ListPveNodeService(c *gin.Context) {
	var res response.PveNode
	req, url, err := pveRequest("/nodes", &res)
	if err != nil {
		sendError(c, err)
	}
	_, err = req.Get(url)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", res.Data)
}
