package frp_admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateFrpAdminToken struct {
	Name string `json:"name"`
	User int    `json:"user"`
}

func (cfat CreateFrpAdminToken) Parse(c *gin.Context) CreateFrpAdminToken {
	if err := c.ShouldBindJSON(&cfat); err != nil {
		response.SendError(c, err)
	}
	return cfat
}

type UpdateFrpAdminToken struct {
	Name   string `json:"name"`
	User   int    `json:"user"`
	Enable int    `json:"enable"`
}

func (ufat UpdateFrpAdminToken) Parse(c *gin.Context) UpdateFrpAdminToken {
	if err := c.ShouldBindJSON(&ufat); err != nil {
		response.SendError(c, err)
	}
	return ufat
}
