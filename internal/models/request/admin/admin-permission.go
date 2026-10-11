package admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateAdminPermission struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}

func (cap CreateAdminPermission) Parse(c *gin.Context) CreateAdminPermission {
	if err := c.ShouldBindJSON(&cap); err != nil {
		response.SendError(c, err)
	}
	return cap
}

type UpdateAdminPermission struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}

func (uap UpdateAdminPermission) Parse(c *gin.Context) UpdateAdminPermission {
	if err := c.ShouldBindJSON(&uap); err != nil {
		response.SendError(c, err)
	}
	return uap
}
