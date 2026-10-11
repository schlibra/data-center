package admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateAdminSettings struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (cas CreateAdminSettings) Parse(c *gin.Context) CreateAdminSettings {
	if err := c.ShouldBindJSON(&cas); err != nil {
		response.SendError(c, err)
	}
	return cas
}

type UpdateAdminSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (uas UpdateAdminSettings) Parse(c *gin.Context) UpdateAdminSettings {
	if err := c.ShouldBindJSON(&uas); err != nil {
		response.SendError(c, err)
	}
	return uas
}
