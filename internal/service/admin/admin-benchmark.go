package admin

import (
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func DoAdminBenchmarkService(c *gin.Context) {
	user := new(repository.User)
	if err := user.New(); err != nil {
		sendError(c, err)
	}
	defer user.Close()
	group := new(repository.Group)
	group.DB = user.DB
	permission := new(repository.Permission)
	permission.DB = user.DB
	settings := new(repository.Settings)
	settings.DB = user.DB
	frpRule := new(repository.FrpRule)
	frpRule.DB = user.DB
	frpToken := new(repository.FrpToken)
	frpToken.DB = user.DB
	mailDomain := new(repository.MailDomain)
	mailDomain.DB = user.DB
	userRow, err := user.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	groupRow, err := group.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	permissionRow, err := permission.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	settingsRow, err := settings.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	frpRuleRow, err := frpRule.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	frpTokenRow, err := frpToken.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	mailDomainRow, err := mailDomain.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", H{
		"users":        userRow,
		"groups":       groupRow,
		"permissions":  permissionRow,
		"settings":     settingsRow,
		"frp_rules":    frpRuleRow,
		"frp_tokens":   frpTokenRow,
		"mail_domains": mailDomainRow,
	})
}
