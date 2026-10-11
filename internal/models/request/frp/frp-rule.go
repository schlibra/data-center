package frp

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type CreateFrpRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
}

func (cfr CreateFrpRule) Parse(c *gin.Context) CreateFrpRule {
	if err := c.ShouldBindJSON(&cfr); err != nil {
		response.SendError(c, err)
	}
	return cfr
}

type UpdateFrpRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
}

func (ufr UpdateFrpRule) Parse(c *gin.Context) UpdateFrpRule {
	if err := c.ShouldBindJSON(&ufr); err != nil {
		response.SendError(c, err)
	}
	return ufr
}
