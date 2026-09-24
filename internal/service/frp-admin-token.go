package service

import (
	"data-center/internal/models/request"
	"data-center/internal/repository"
	"data-center/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ListFrpAdminTokenService(c *gin.Context) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	tokens, err := frpToken.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "frp.token.got", tokens)
}
func GetFrpAdminTokenService(c *gin.Context, id int) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	token, err := frpToken.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
	}
	sendI18n(c, 200, "frp.token.got", token)
}
func CreateFrpAdminTokenService(c *gin.Context, req request.CreateFrpAdminToken) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	_, err = frpToken.SelectByName(req.Name)
	if err == nil {
		sendI18n(c, 400, "frp.token.exist", nil)
	}
	_, err = user.SelectById(req.User)
	if err != nil {
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
func UpdateFrpAdminTokenService(c *gin.Context, id int, req request.UpdateFrpAdminToken) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	tokenRow, err := frpToken.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	if _, err := frpToken.SelectByName(req.Name); err == nil {
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
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	if _, err := frpToken.SelectById(id); err != nil {
		sendI18n(c, 400, "frp.token.not_exist", H{
			"error": err.Error(),
		})
		return
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
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
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
