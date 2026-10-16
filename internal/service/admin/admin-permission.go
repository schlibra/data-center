package admin

import (
	"data-center/internal/models"
	"data-center/internal/models/request/admin"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminPermissionService(c *gin.Context) {
	permission := new(repository.Permission)
	if err := permission.New(); err != nil {
		sendError(c, err)
	}
	defer permission.Close()
	permissions, err := permission.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	permissionList := make(map[int]*models.PermissionTable)
	for _, permission := range permissions {
		if permission.Parent == 0 {
			permissionList[permission.ID] = &permission
		}
	}
	for _, permission := range permissions {
		if permission.Parent != 0 {
			permissionList[permission.Parent].Children = append(permissionList[permission.Parent].Children, permission)
		}
	}
	permissionResult := make([]models.PermissionTable, 0)
	for _, permission := range permissionList {
		if permission.Children == nil {
			permission.Children = make([]models.PermissionTable, 0)
		}
		permissionResult = append(permissionResult, *permission)
	}
	sendI18n(c, 200, "admin.permission.get_success", permissionResult)
}
func GetAdminPermissionService(c *gin.Context, id int) {
	permission := new(repository.Permission)
	if err := permission.New(); err != nil {
		sendError(c, err)
	}
	defer permission.Close()
	permissionRow, err := permission.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "admin.permission.get_success", permissionRow)
}
func CreateAdminPermissionService(c *gin.Context, req admin.CreateAdminPermission) {
	permission := new(repository.Permission)
	if err := permission.New(); err != nil {
		sendError(c, err)
	}
	defer permission.Close()
	if _, err := permission.SelectByKey(req.Key); err == nil {
		sendI18n(c, 400, "admin.permission.exist", req.Key)
	}
	if nw(c) {
		if _, err := permission.Insert(req.Key, req.Name, req.Parent); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.permission.create_success", req.Key)
		}
	}
}
func UpdateAdminPermissionService(c *gin.Context, id int, req admin.UpdateAdminPermission) {
	permission := new(repository.Permission)
	if err := permission.New(); err != nil {
		sendError(c, err)
	}
	defer permission.Close()
	permissionRow, err := permission.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "admin.permission.not_exists", nil)
		return
	}
	if _, err := permission.SelectByKey(req.Key); err == nil && permissionRow.Key != req.Key {
		sendI18n(c, 400, "admin.permission.exist", nil)
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
	permission := new(repository.Permission)
	if err := permission.New(); err != nil {
		sendError(c, err)
	}
	defer permission.Close()
	if _, err := permission.SelectById(id); err != nil {
		sendI18n(c, 400, "admin.permission.not_exists", nil)
	}
	childPermissions, err := permission.SelectByParent(id)
	if err != nil {
		sendError(c, err)
	}
	if len(childPermissions) > 0 {
		sendI18n(c, 400, "admin.permission.has_child", nil)
	}
	if nw(c) {
		if _, err := permission.DeleteById(id); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "admin.permission.delete_success", nil)
		}
	}
}
