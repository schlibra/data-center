package frp_admin

import (
	"data-center/internal/models/request/frp_admin"
	"data-center/internal/repository"
	"data-center/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ListFrpAdminTokenService(c *gin.Context) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	tokens, err := frpToken.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.token.got", tokens)
}
func GetFrpAdminTokenService(c *gin.Context, id int) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
	}
	sendI18n(c, 200, "frp.token.got", token)
}
func CreateFrpAdminTokenService(c *gin.Context, req frp_admin.CreateFrpAdminToken) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	user := new(repository.User)
	user.DB = frpToken.DB
	if _, err := frpToken.SelectByName(req.Name); err == nil {
		sendI18n(c, 400, "frp.token.exist", nil)
	}
	if _, err := user.SelectById(req.User); err != nil {
		sendI18n(c, 400, "admin.user.not_exist", H{
			"error": err.Error(),
		})
	}
	rndToken, err := utils.GenerateRandomString(9)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		if _, err := frpToken.Insert(req.Name, rndToken, req.User, 1); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "frp.token.created", nil)
		}
	}
}
func UpdateFrpAdminTokenService(c *gin.Context, id int, req frp_admin.UpdateFrpAdminToken) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	user := new(repository.User)
	user.DB = frpToken.DB
	tokenRow, err := frpToken.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	if _, err := frpToken.SelectByName(req.Name); err == nil && req.Name != tokenRow.Name {
		sendI18n(c, 400, "frp.token.exist", nil)
	}
	if _, err := user.SelectById(req.User); err != nil {
		sendI18n(c, 400, "admin.user.not_exist", H{
			"error": err.Error(),
		})
	}
	if nw(c) {
		if _, err := frpToken.UpdateById(id, req.Name, tokenRow.Token, req.User, req.Enable); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "frp.token.updated", nil)
		}
	}
}
func DeleteFrpAdminTokenService(c *gin.Context, id int) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	frpRule := new(repository.FrpRule)
	frpRule.DB = frpToken.DB
	if _, err := frpToken.SelectById(id); err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	rules, err := frpRule.SelectByToken(id)
	if err != nil {
		sendError(c, err)
	}
	if len(rules) > 0 {
		sendI18n(c, 400, "frp.token.has_rule", nil)
	}
	if nw(c) {
		if _, err := frpToken.DeleteById(id); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "frp.token.deleted", nil)
		}
	}
}
func GenerateFrpAdminTokenService(c *gin.Context, id int) {
	frpToken := new(repository.FrpToken)
	if err := frpToken.New(); err != nil {
		sendError(c, err)
	}
	defer frpToken.Close()
	tokenRow, err := frpToken.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	rndToken, err := utils.GenerateRandomString(9)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		if _, err := frpToken.UpdateById(id, tokenRow.Name, rndToken, tokenRow.User, tokenRow.Enable); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "frp.token.generated", nil)
		}
	}
}
