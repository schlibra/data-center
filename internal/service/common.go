package service

import (
	"data-center/pkg/response"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type H map[string]any

func sendJson(c *gin.Context, code int, message string, data any) {
	response.SendJson(c, code, message, data)
}
func sendI18n(c *gin.Context, code int, messageId string, data any) {
	response.SendI18n(c, code, messageId, data)
}
func sendError(c *gin.Context, err error) {
	response.SendError(c, err)
}
func nw(c *gin.Context) bool {
	return response.NW(c)
}

func closeDB(db *sql.DB) {
	_ = db.Close()
}
