package mail

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateMailList struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cml CreateMailList) Parse(c *gin.Context) CreateMailList {
	if err := c.ShouldBindJSON(&cml); err != nil {
		response.SendError(c, err)
	}
	return cml
}

type UpdateMailList struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (uml UpdateMailList) Parse(c *gin.Context) UpdateMailList {
	if err := c.ShouldBindJSON(&uml); err != nil {
		response.SendError(c, err)
	}
	return uml
}
