package service

import (
	"data-center/internal/models/request"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminGroupService(c *gin.Context) {
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(group.DB)
	groups, err := group.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "admin.group.get_success", groups)
}
func GetAdminGroupService(c *gin.Context, id int) {
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(group.DB)
	groupRow, err := group.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.group.not_exist", H{
			"error": err.Error(),
		})
	}
	sendI18n(c, 200, "admin.group.get_success", groupRow)
}
func CreateAdminGroupService(c *gin.Context, req request.CreateAdminGroup) {
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(group.DB)
	if _, err := group.SelectByName(req.Name); err == nil {
		sendI18n(c, 400, "admin.group.exist", nil)
	}
	if nw(c) {
		if _, err := group.Insert(req.Name, req.Admin, req.Permission); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.group.create_success", nil)
		}
	}
}
func UpdateAdminGroupService(c *gin.Context, id int, req request.UpdateAdminGroup) {
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(group.DB)
	groupRow, err := group.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.group.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	if _, err := group.SelectByName(req.Name); err != nil {
		sendI18n(c, 400, "admin.group.not_exist", nil)
	}
	if _, err := group.SelectByName(req.Name); err == nil && groupRow.Name != req.Name {
		sendI18n(c, 400, "admin.group.exist", nil)
	}
	if nw(c) {
		if _, err := group.UpdateById(id, req.Name, req.Admin, req.Permission); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.group.update_success", nil)
		}
	}
}
func DeleteAdminGroupService(c *gin.Context, id int) {
	group, err := repository.NewGroup()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(group.DB)
	if _, err := group.SelectById(id); err != nil {
		sendI18n(c, 400, "admin.group.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	if nw(c) {
		if _, err := group.DeleteById(id); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.group.delete_success", nil)
		}
	}
}
