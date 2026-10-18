package pve

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type SetPveQemuStatus struct {
	Action string `json:"action"`
}

func (spqs SetPveQemuStatus) Parse(c *gin.Context) SetPveQemuStatus {
	if err := c.ShouldBindJSON(&spqs); err != nil {
		response.SendError(c, err)
	}
	return spqs
}
