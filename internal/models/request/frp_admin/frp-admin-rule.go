package frp_admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateFrpAdminRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
	User  int `json:"user"`
}

func (cfar CreateFrpAdminRule) Parse(c *gin.Context) CreateFrpAdminRule {
	if err := c.ShouldBindJSON(&cfar); err != nil {
		response.SendError(c, err)
	}
	return cfar
}

type UpdateFrpAdminRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
	User  int `json:"user"`
}

func (ufar UpdateFrpAdminRule) Parse(c *gin.Context) UpdateFrpAdminRule {
	if err := c.ShouldBindJSON(&ufar); err != nil {
		response.SendError(c, err)
	}
	return ufar
}
