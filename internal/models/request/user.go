package request

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ul UserLogin) Parse(c *gin.Context) UserLogin {
	if err := c.ShouldBindJSON(&ul); err != nil {
		response.SendError(c, err)
	}
	return ul
}

type UserLoginKey struct {
	Username string `json:"username"`
}

func (ulk UserLoginKey) Parse(c *gin.Context) UserLoginKey {
	if err := c.ShouldBindJSON(&ulk); err != nil {
		response.SendError(c, err)
	}
	return ulk
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func (ur UserRegister) Parse(c *gin.Context) UserRegister {
	if err := c.ShouldBindJSON(&ur); err != nil {
		response.SendError(c, err)
	}
	return ur
}

type UserRegisterKey struct {
	Username string `json:"username"`
}

func (urk UserRegisterKey) Parse(c *gin.Context) UserRegisterKey {
	if err := c.ShouldBindJSON(&urk); err != nil {
		response.SendError(c, err)
	}
	return urk
}

type UserUpdate struct {
	Nickname string `json:"nickname"`
}

func (uu UserUpdate) Parse(c *gin.Context) UserUpdate {
	if err := c.ShouldBindJSON(&uu); err != nil {
		response.SendError(c, err)
	}
	return uu
}

type UserPassword struct {
	Password string `json:"password"`
}

func (up UserPassword) Parse(c *gin.Context) UserPassword {
	if err := c.ShouldBindJSON(&up); err != nil {
		response.SendError(c, err)
	}
	return up
}
