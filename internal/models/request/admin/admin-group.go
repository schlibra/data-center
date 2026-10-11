package admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateAdminGroup struct {
	Name       string `json:"name"`
	Admin      int    `json:"admin"`
	Permission string `json:"permission"`
}

func (cag CreateAdminGroup) Parse(c *gin.Context) CreateAdminGroup {
	if err := c.ShouldBindJSON(&cag); err != nil {
		response.SendError(c, err)
	}
	return cag
}

type UpdateAdminGroup struct {
	Name       string `json:"name"`
	Admin      int    `json:"admin"`
	Permission string `json:"permission"`
}

func (uag UpdateAdminGroup) Parse(c *gin.Context) UpdateAdminGroup {
	if err := c.ShouldBindJSON(&uag); err != nil {
		response.SendError(c, err)
	}
	return uag
}
