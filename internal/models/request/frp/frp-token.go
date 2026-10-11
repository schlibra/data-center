package frp

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateFrpToken struct {
	Name string `json:"name"`
}

func (cft CreateFrpToken) Parse(c *gin.Context) CreateFrpToken {
	if err := c.ShouldBindBodyWithJSON(&cft); err != nil {
		response.SendError(c, err)
	}
	return cft
}

type UpdateFrpToken struct {
	Name   string `json:"name"`
	Enable int    `json:"enable"`
}
