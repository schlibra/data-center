package service

import (
	"data-center/internal/models"
	"data-center/internal/models/request"
	"data-center/internal/repository"
	"data-center/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateFrpTokenService(c *gin.Context, req request.FrpTokenCreate, row models.UserTable) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	if _, err := frpToken.SelectByName(req.Name); err == nil {
		sendI18n(c, 400, "frp.token.exist", nil)
	}
	token, err := utils.GenerateRandomString(9)
	if err != nil {
		sendError(c, err)
	}
	if !c.Writer.Written() {
		if _, err := frpToken.Insert(req.Name, token, row.ID, 1); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.token.created", nil)
	}
}

func ListFrpTokenService(c *gin.Context, row models.UserTable) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	tokens, err := frpToken.SelectByUser(row.ID)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.token.got", tokens)
}

func GetFrpTokenService(c *gin.Context, row models.UserTable, id int) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	if token.User != row.ID {
		sendI18n(c, 403, "frp.token.forbidden", nil)
	}
	sendI18n(c, 200, "frp.token.got", H{
		"token": token,
	})
}

func UpdateFrpTokenService(c *gin.Context, req request.FrpTokenUpdate, row models.UserTable, id int) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	if token.User != row.ID {
		sendI18n(c, 403, "frp.token.forbidden", nil)
	}
	if _, err := frpToken.SelectByName(req.Name); err == nil && req.Name != token.Name {
		sendI18n(c, 400, "frp.token.exist", nil)
	}
	if !c.Writer.Written() {
		_, err := frpToken.UpdateById(token.ID, req.Name, token.Token, row.ID, req.Enable)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.token.updated", nil)
	}
}

func DeleteFrpTokenService(c *gin.Context, row models.UserTable, id int) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpRule.DB)
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	if token.User != row.ID {
		sendI18n(c, 403, "frp.token.forbidden", nil)
	}
	rules, err := frpRule.SelectByToken(id)
	if err != nil {
		sendError(c, err)
	}
	if len(rules) > 0 {
		sendI18n(c, 400, "frp.token.has_rule", nil)
	}
	if !c.Writer.Written() {
		_, err := frpToken.DeleteById(token.ID)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.token.deleted", nil)
	}
}

func GenerateFrpTokenService(c *gin.Context, row models.UserTable, id int) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	if token.User != row.ID {
		sendI18n(c, 403, "frp.token.forbidden", nil)
	}
	if !c.Writer.Written() {
		t, err := utils.GenerateRandomString(9)
		if err != nil {
			sendError(c, err)
		}
		_, err = frpToken.UpdateById(token.ID, token.Name, t, row.ID, 1)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "frp.token.generated", nil)
	}
}
