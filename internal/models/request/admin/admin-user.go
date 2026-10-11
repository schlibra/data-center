package admin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateAdminUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Group    int    `json:"group"`
	Enable   int    `json:"enable"`
}

func (cau CreateAdminUser) Parse(c *gin.Context) CreateAdminUser {
	if err := c.ShouldBindJSON(&cau); err != nil {
		response.SendError(c, err)
	}
	return cau
}

type UpdateAdminUser struct {
	Nickname string `json:"nickname"`
	Group    int    `json:"group"`
	Enable   int    `json:"enable"`
}

func (uau UpdateAdminUser) Parse(c *gin.Context) UpdateAdminUser {
	if err := c.ShouldBindJSON(&uau); err != nil {
		response.SendError(c, err)
	}
	return uau
}

type PasswordAdminUser struct {
	Password string `json:"password"`
}

func (pau PasswordAdminUser) Parse(c *gin.Context) PasswordAdminUser {
	if err := c.ShouldBindJSON(&pau); err != nil {
		response.SendError(c, err)
	}
	return pau
}
