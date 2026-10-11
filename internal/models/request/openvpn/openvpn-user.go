package openvpn

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type createOpenVPNUserSrcAddr struct {
	Object []struct {
		Type   int    `json:"type"`
		GPName string `json:"gp_name"`
		GID    string `json:"gid"`
	} `json:"object"`
}

type CreateOpenVPNUser struct {
	Username string                   `json:"username"`
	Password string                   `json:"password"`
	Enabled  string                   `json:"enabled"`
	Share    int                      `json:"share"`
	SrcAddr  createOpenVPNUserSrcAddr `json:"src_addr"`
}

func (covu CreateOpenVPNUser) Parse(c *gin.Context) CreateOpenVPNUser {
	if err := c.ShouldBindJSON(&covu); err != nil {
		response.SendError(c, err)
	}
	return covu
}

type UpdateOpenVPNUser struct {
	Username string                   `json:"username"`
	Password string                   `json:"password"`
	Enabled  string                   `json:"enabled"`
	Share    int                      `json:"share"`
	SrcAddr  createOpenVPNUserSrcAddr `json:"src_addr"`
}

func (uou UpdateOpenVPNUser) Parse(c *gin.Context) UpdateOpenVPNUser {
	if err := c.ShouldBindJSON(&uou); err != nil {
		response.SendError(c, err)
	}
	return uou
}
