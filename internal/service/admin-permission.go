package service

import (
	"data-center/internal/models/request"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminPermissionService(c *gin.Context) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	permissions, err := permission.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "admin.permission.get_success", permissions)
}
func GetAdminPermissionService(c *gin.Context, id int) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	permissionRow, err := permission.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "admin.permission.get_success", permissionRow)
}
func CreateAdminPermissionService(c *gin.Context, req request.CreateAdminPermission) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	if _, err := permission.SelectByKey(req.Key); err == nil {
		sendI18n(c, 400, "admin.permission.key_exists", req.Key)
	}
	if nw(c) {
		if _, err := permission.Insert(req.Key, req.Name, req.Parent); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.permission.create_success", req.Key)
		}
	}
}
func UpdateAdminPermissionService(c *gin.Context, id int, req request.UpdateAdminPermission) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	if _, err := permission.SelectById(id); err == nil {
		sendI18n(c, 400, "admin.permission.exists", nil)
	}
	if nw(c) {
		if _, err := permission.UpdateByID(id, req.Key, req.Name, req.Parent); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.permission.update_success", req.Key)
		}
	}
}
func DeleteAdminPermissionService(c *gin.Context, id int) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	if _, err := permission.SelectById(id); err == nil {
		sendI18n(c, 400, "admin.permission.exists", nil)
	}
	if nw(c) {
		if _, err := permission.DeleteById(id); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.permission.delete_success", nil)
		}
	}
}
