package handler

import (
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

// GetConfigHandler 获取配置
// @Summary 获取配置
// @Tags 配置
// @Param key path string true "配置键"
// @Success 200 {object} response.Response
// @Router /config/{key} [get]
func GetConfigHandler(c *gin.Context) {
	type KeyRequest struct {
		Key string `uri:"key"`
	}
	var req KeyRequest
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GetConfigService(c, req.Key)
}
