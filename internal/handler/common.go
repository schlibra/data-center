package handler

import (
	"data-center/internal/models"
	"data-center/internal/repository"
	"data-center/pkg/response"
	"data-center/pkg/utils"
	"database/sql"
	"encoding/json"
	"slices"
	"strings"

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
func ParseToken(c *gin.Context) models.UserTable {
	return parseToken(c)
}
func parseToken(c *gin.Context) models.UserTable {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}
	data, err := utils.JwtUserCheck(token)
	if err != nil {
		sendI18n(c, 401, "user.token.invalid", H{
			"error": err.Error(),
		})
		return models.UserTable{}
	}
	row, err := user.SelectByUsername(data.Username)
	if err != nil {
		sendError(c, err)
	}
	if row.ID != data.UserID {
		sendI18n(c, 401, "user.token.id_not_match", nil)
	}
	if row.TokenID != data.TokenID && row.ApiId != data.TokenID {
		sendI18n(c, 401, "user.token.invalid", nil)
	}
	return row
}
func CheckPermission(c *gin.Context, pmsKey string) models.UserTable {
	return checkPermission(c, pmsKey)
}
func checkPermission(c *gin.Context, pmsKey string) models.UserTable {
	row := parseToken(c)
	if row.GroupInfo.Admin == 1 {
		return row
	}
	var groupPmsList []int
	if err := json.Unmarshal([]byte(row.GroupInfo.Permission), &groupPmsList); err != nil {
		sendError(c, err)
	}
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	p, err := permission.SelectByKey(pmsKey)
	if err != nil {
		sendI18n(c, 403, "user.permission.denied", H{
			"error": err.Error(),
		})
	}
	if !slices.Contains(groupPmsList, p.ID) {
		sendI18n(c, 403, "user.permission.denied", nil)
	}
	return row
}
func CheckAdmin(c *gin.Context) models.UserTable {
	return checkAdmin(c)
}
func checkAdmin(c *gin.Context) models.UserTable {
	row := parseToken(c)
	if row.GroupInfo.Admin == 0 {
		sendI18n(c, 403, "user.admin.no_permission", nil)
	}
	return row
}
func PortCheck(c *gin.Context, min int, max int) {
	portCheck(c, min, max)
}
func portCheck(c *gin.Context, min int, max int) {
	if min < 1 || max < 1 || min > 65535 || max > 65535 || min > max {
		sendI18n(c, 400, "frp.rule.port_not_valid", nil)
	}
}
