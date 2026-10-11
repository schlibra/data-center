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

func (uft UpdateFrpToken) Parse(c *gin.Context) UpdateFrpToken {
	if err := c.ShouldBindBodyWithJSON(&uft); err != nil {
		response.SendError(c, err)
	}
	return uft
}
