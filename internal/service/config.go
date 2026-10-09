package service

import (
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetConfigService(c *gin.Context, key string) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	row, err := settings.SelectByKey(key)
	if err != nil {
		sendJson(c, 200, "success", "")
	} else {
		sendJson(c, 200, "success", row.Value)
	}
}
