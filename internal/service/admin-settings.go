package service

import (
	"data-center/internal/models/request"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminSettingsService(c *gin.Context) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	rows, err := settings.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "settings.get_success", rows)
}
func GetAdminSettingsService(c *gin.Context, id int) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	row, err := settings.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "settings.not_exists", nil)
	}
	sendI18n(c, 200, "settings.get_success", row)
}
func CreateAdminSettingsService(c *gin.Context, req request.CreateAdminSettings) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	_, err = settings.SelectByKey(req.Key)
	if err == nil {
		sendI18n(c, 400, "settings.exists", nil)
	}
	if nw(c) {
		_, err := settings.Insert(req.Key, req.Name, req.Value)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "settings.create_success", nil)
	}
}
func UpdateAdminSettingsService(c *gin.Context, id int, req request.UpdateAdminSettings) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	row, err := settings.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "settings.not_exists", nil)
		return
	}
	if nw(c) {
		_, err := settings.UpdateById(id, row.Key, req.Name, req.Value)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "settings.update_success", nil)
	}
}
func DeleteAdminSettingsService(c *gin.Context, id int) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	_, err = settings.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "settings.not_exists", nil)
		return
	}
	if nw(c) {
		_, err := settings.DeleteById(id)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "settings.delete_success", nil)
	}
}
func SetAdminSettingsService(c *gin.Context, req request.CreateAdminSettings) {
	settings, err := repository.NewSettings()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(settings.DB)
	row, err := settings.SelectByKey(req.Key)
	if err != nil {
		_, err = settings.Insert(req.Key, req.Name, req.Value)
	} else {
		_, err = settings.UpdateById(row.ID, row.Key, req.Name, req.Value)
	}
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "settings.set_success", nil)
}
