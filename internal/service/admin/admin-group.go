package admin

import (
	"data-center/internal/models/request/admin"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminGroupService(c *gin.Context) {
	group := new(repository.Group)
	if err := group.New(); err != nil {
		sendError(c, err)
	}
	defer group.Close()
	groups, err := group.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "admin.group.get_success", groups)
}
func GetAdminGroupService(c *gin.Context, id int) {
	group := new(repository.Group)
	if err := group.New(); err != nil {
		sendError(c, err)
	}
	defer group.Close()
	groupRow, err := group.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.group.not_exist", H{
			"error": err.Error(),
		})
	}
	sendI18n(c, 200, "admin.group.get_success", groupRow)
}
func CreateAdminGroupService(c *gin.Context, req admin.CreateAdminGroup) {
	group := new(repository.Group)
	if err := group.New(); err != nil {
		sendError(c, err)
	}
	defer group.Close()
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
func UpdateAdminGroupService(c *gin.Context, id int, req admin.UpdateAdminGroup) {
	group := new(repository.Group)
	if err := group.New(); err != nil {
		sendError(c, err)
	}
	defer group.Close()
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
	group := new(repository.Group)
	if err := group.New(); err != nil {
		sendError(c, err)
	}
	defer group.Close()
	user := new(repository.User)
	if err := user.New(); err != nil {
		sendError(c, err)
	}
	defer user.Close()
	if _, err := group.SelectById(id); err != nil {
		sendI18n(c, 400, "admin.group.not_exist", H{
			"error": err.Error(),
		})
		return
	}
	users, err := user.SelectByGroup(id)
	if err != nil {
		sendError(c, err)
	}
	if len(users) > 0 {
		sendI18n(c, 400, "admin.group.has_user", nil)
	}
	if nw(c) {
		if _, err := group.DeleteById(id); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.group.delete_success", nil)
		}
	}
}
