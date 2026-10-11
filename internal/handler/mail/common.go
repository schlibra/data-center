package mail

import (
	"data-center/internal/handler"
	"data-center/internal/models"
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
func closeDB(db *sql.DB) {
	_ = db.Close()
}
func parseToken(c *gin.Context) models.UserTable {
	return handler.ParseToken(c)
}
func checkPermission(c *gin.Context, pmsKey string) models.UserTable {
	return handler.CheckPermission(c, pmsKey)
}
func checkAdmin(c *gin.Context) models.UserTable {
	return handler.CheckAdmin(c)
}
func portCheck(c *gin.Context, min int, max int) {
	handler.PortCheck(c, min, max)
}
func getId(c *gin.Context) int {
	return handler.GetId(c)
}
func getDomain(c *gin.Context) int {
	type ReqDomain struct {
		Domain int `uri:"domain"`
	}
	var reqDomain ReqDomain
	if err := c.ShouldBindUri(&reqDomain); err != nil {
		sendError(c, err)
	}
	return reqDomain.Domain
}
func getRecord(c *gin.Context) string {
	type ReqRecord struct {
		Record string `uri:"record"`
	}
	var reqRecord ReqRecord
	if err := c.ShouldBindUri(&reqRecord); err != nil {
		sendError(c, err)
	}
	return reqRecord.Record
}
