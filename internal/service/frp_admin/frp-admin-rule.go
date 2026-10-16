package frp_admin

import (
	"data-center/internal/models/request/frp_admin"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListFrpAdminRuleService(c *gin.Context) {
	frpRule := new(repository.FrpRule)
	if err := frpRule.New(); err != nil {
		sendError(c, err)
	}
	defer frpRule.Close()
	rules, err := frpRule.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.rule.get_success", rules)
}
func GetFrpAdminRuleService(c *gin.Context, id int) {
	frpRule := new(repository.FrpRule)
	if err := frpRule.New(); err != nil {
		sendError(c, err)
	}
	defer frpRule.Close()
	rule, err := frpRule.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.rule.get_success", rule)
}
func CreateFrpAdminRuleService(c *gin.Context, req frp_admin.CreateFrpAdminRule) {
	frpRule := new(repository.FrpRule)
	if err := frpRule.New(); err != nil {
		sendError(c, err)
	}
	defer frpRule.Close()
	frpToken := new(repository.FrpToken)
	frpToken.DB = frpRule.DB
	user := new(repository.User)
	user.DB = frpRule.DB
	if _, err := frpToken.SelectById(req.Token); err != nil {
		sendI18n(c, 400, "frp.rule.token_not_exist", nil)
	}
	if _, err := user.SelectById(req.User); err != nil {
		sendI18n(c, 400, "frp.rule.user_not_exist", nil)
	}
	if nw(c) {
		if _, err := frpRule.Insert(req.Min, req.Max, req.Token, req.User); err != nil {
			sendI18n(c, 400, "frp.rule.create_failed", nil)
		}
		sendI18n(c, 200, "frp.rule.create_success", nil)
	}
}
func UpdateFrpAdminRuleService(c *gin.Context, req frp_admin.UpdateFrpAdminRule, id int) {
	frpRule := new(repository.FrpRule)
	if err := frpRule.New(); err != nil {
		sendError(c, err)
	}
	defer frpRule.Close()
	frpToken := new(repository.FrpToken)
	frpToken.DB = frpRule.DB
	user := new(repository.User)
	user.DB = frpRule.DB
	if _, err := frpToken.SelectById(req.Token); err != nil {
		sendI18n(c, 400, "frp.rule.token_not_exist", nil)
	}
	if _, err := user.SelectById(req.User); err != nil {
		sendI18n(c, 400, "frp.rule.user_not_exist", nil)
	}
	if _, err := frpRule.SelectById(id); err != nil {
		sendI18n(c, 400, "frp.rule.not_exist", nil)
	}
	if nw(c) {
		if _, err := frpRule.UpdateById(id, req.Min, req.Max, req.Token, req.User); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.rule.update_success", nil)
	}
}
func DeleteFrpAdminRuleService(c *gin.Context, id int) {
	frpRule := new(repository.FrpRule)
	if err := frpRule.New(); err != nil {
		sendError(c, err)
	}
	defer frpRule.Close()
	if _, err := frpRule.SelectById(id); err != nil {
		sendI18n(c, 400, "frp.rule.not_exist", nil)
	}
	if nw(c) {
		if _, err := frpRule.DeleteById(id); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.rule.delete_success", nil)
	}
}
