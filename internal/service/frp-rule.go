package service

import (
	"data-center/internal/models"
	"data-center/internal/models/request"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func CreateFrpRuleService(c *gin.Context, row models.UserTable, req request.CreateFrpRule) {
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	_, err = frpToken.SelectById(req.Token)
	if err != nil {
		sendI18n(c, 400, "frp.rule.token_not_exist", nil)
	}
	if nw(c) {
		if _, err := frpRule.Insert(req.Min, req.Max, req.Token, row.ID); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.rule.create_success", nil)
	}

}
func ListFrpRuleService(c *gin.Context, row models.UserTable) {
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	rules, err := frpRule.SelectByUser(row.ID)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.rule.get_success", rules)
}
func GetFrpRuleService(c *gin.Context, row models.UserTable, id int) {
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	rule, err := frpRule.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.rule.not_exist", nil)
		return
	}
	if row.ID != rule.User {
		sendI18n(c, 400, "frp.rule.forbidden", nil)
	}
	sendI18n(c, 200, "frp.rule.get_success", rule)
}
func UpdateFrpRuleService(c *gin.Context, row models.UserTable, req request.UpdateFrpRule, id int) {
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	rule, err := frpRule.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.rule.not_exist", nil)
		return
	}
	if row.ID != rule.User {
		sendI18n(c, 400, "frp.rule.forbidden", nil)
	}
	if nw(c) {
		if _, err := frpRule.UpdateById(id, req.Min, req.Max, req.Token, row.ID); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.rule.update_success", nil)
	}
}
func DeleteFrpRuleService(c *gin.Context, row models.UserTable, id int) {
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	rule, err := frpRule.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.rule.not_exist", nil)
		return
	}
	if row.ID != rule.User {
		sendI18n(c, 400, "frp.rule.forbidden", nil)
	}
	if nw(c) {
		if _, err := frpRule.DeleteById(id); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.rule.delete_success", nil)
	}
}
