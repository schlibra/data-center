package mail

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateMailDomain struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	TableId string `json:"tableId"`
	Server  string `json:"server"`
	Port    int    `json:"port"`
	SSL     string `json:"ssl"`
}

func (cmd CreateMailDomain) Parse(c *gin.Context) CreateMailDomain {
	if err := c.ShouldBindJSON(&cmd); err != nil {
		response.SendError(c, err)
	}
	return cmd
}

type UpdateMailDomain struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	TableId string `json:"tableId"`
	Server  string `json:"server"`
	Port    int    `json:"port"`
	SSL     string `json:"ssl"`
}

func (cmd UpdateMailDomain) Parse(c *gin.Context) UpdateMailDomain {
	if err := c.ShouldBindJSON(&cmd); err != nil {
		response.SendError(c, err)
	}
	return cmd
}
