package openvpn

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type openVPNGroupValue struct {
	IP      string `json:"ip"`
	Comment string `json:"comment"`
}
type CreateOpenVPNGroup struct {
	Name  string              `json:"name"`
	Value []openVPNGroupValue `json:"value"`
}

func (covg CreateOpenVPNGroup) Parse(c *gin.Context) CreateOpenVPNGroup {
	if err := c.ShouldBindJSON(&covg); err != nil {
		response.SendError(c, err)
	}
	return covg
}

type UpdateOpenVPNGroup struct {
	Name  string              `json:"name"`
	Value []openVPNGroupValue `json:"value"`
}

func (uovg UpdateOpenVPNGroup) Parse(c *gin.Context) UpdateOpenVPNGroup {
	if err := c.ShouldBindJSON(&uovg); err != nil {
		response.SendError(c, err)
	}
	return uovg
}
