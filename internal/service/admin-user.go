package service

import (
	"data-center/internal/models"
	"data-center/internal/models/request"
	"data-center/internal/repository"
	"data-center/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ListAdminUserService(c *gin.Context) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	users, err := user.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	var usersList []models.UserTable
	for _, user := range users {
		user.Password = "********"
		user.TokenID = "********"
		usersList = append(usersList, user)
	}
	sendI18n(c, 200, "admin.user.get_success", usersList)
}
func GetAdminUserService(c *gin.Context, id int) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	userRow, err := user.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.user.not_exist", nil)
		return
	}
	userRow.Password = "********"
	userRow.TokenID = "********"
	sendI18n(c, 200, "admin.user.get_success", userRow)
}
func CreateAdminUserService(c *gin.Context, req request.CreateAdminUser, id int) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	if _, err := group.SelectById(req.Group); err != nil {
		sendI18n(c, 400, "admin.user.group_not_exist", nil)
	}
	if _, err := user.SelectById(id); err == nil {
		sendI18n(c, 400, "admin.user.exist", nil)
	}
	password, err := utils.PasswordHash(req.Password)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		if _, err := user.Insert(req.Username, password, req.Nickname, req.Group, req.Enable); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "admin.user.create_success", nil)
	}
}
func UpdateAdminUserService(c *gin.Context, req request.UpdateAdminUser, id int) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	userRow, err := user.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.user.not_exist", nil)
	}
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	if _, err := group.SelectById(req.Group); err != nil {
		sendI18n(c, 400, "admin.user.group_not_exist", nil)
	}
	if nw(c) {
		if _, err := user.UpdateById(id, userRow.Password, req.Nickname, req.Group, req.Enable, userRow.TokenID); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "admin.user.update_success", nil)
	}
}
func DeleteAdminUserService(c *gin.Context, row models.UserTable, id int) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	userRow, err := user.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.user.not_exist", nil)
		return
	}
	if userRow.ID == row.ID {
		sendI18n(c, 403, "admin.user.delete_self", nil)
	}
	if nw(c) {
		if _, err := user.DeleteById(id); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "admin.user.delete_success", nil)
	}
}
func PasswordAdminUserService(c *gin.Context, req request.PasswordAdminUser, id int) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	userRow, err := user.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.user.not_exist", nil)
		return
	}
	password, err := utils.PasswordHash(req.Password)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		if _, err := user.UpdateById(id, password, userRow.Nickname, userRow.Group, userRow.Enable, ""); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "admin.user.password_update_success", nil)
	}
}
