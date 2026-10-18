package pve

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type SetPveLxcStatus struct {
	Action string `json:"action"`
}

func (spls SetPveLxcStatus) Parse(c *gin.Context) SetPveLxcStatus {
	if err := c.ShouldBindJSON(&spls); err != nil {
		response.SendError(c, err)
	}
	return spls
}
