package autologin

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateAdminAutoLogin struct {
	Url      string `json:"url"`
	Site     string `json:"site"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (caal CreateAdminAutoLogin) Parse(c *gin.Context) CreateAdminAutoLogin {
	if err := c.ShouldBindJSON(&caal); err != nil {
		response.SendError(c, err)
	}
	return caal
}

type UpdateAdminAutoLogin struct {
	Url      string `json:"url"`
	Site     string `json:"site"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uaal UpdateAdminAutoLogin) Parse(c *gin.Context) UpdateAdminAutoLogin {
	if err := c.ShouldBindJSON(&uaal); err != nil {
		response.SendError(c, err)
	}
	return uaal
}
