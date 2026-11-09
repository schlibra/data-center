package autologin

import (
	"data-center/internal/models/request/autologin"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListAdminAutoLoginService(c *gin.Context) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	rows, err := autoLogin.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", rows)
}
func GetAdminAutoLoginService(c *gin.Context, id int) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	row, err := autoLogin.SelectById(id)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", row)
}
func CreateAdminAutoLoginService(c *gin.Context, caal autologin.CreateAdminAutoLogin) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	if _, err := autoLogin.Insert(caal.Url, caal.Site, caal.Icon, caal.Name, caal.Username, caal.Password); err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", nil)
}
func UpdateAdminAutoLoginService(c *gin.Context, id int, uaal autologin.UpdateAdminAutoLogin) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	if _, err := autoLogin.SelectById(id); err != nil {
		sendJson(c, 400, "auto login not found", H{
			"error": err.Error(),
		})
	}
	if nw(c) {
		if _, err := autoLogin.UpdateById(id, uaal.Url, uaal.Site, uaal.Icon, uaal.Name, uaal.Username, uaal.Password); err != nil {
			sendError(c, err)
		} else {
			sendJson(c, 200, "success", nil)
		}
	}
}
func DeleteAdminAutoLoginService(c *gin.Context, id int) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	if _, err := autoLogin.SelectById(id); err != nil {
		sendJson(c, 400, "auto login not found", H{
			"error": err.Error(),
		})
	}
	if nw(c) {
		if _, err := autoLogin.DeleteById(id); err != nil {
			sendError(c, err)
		}
		sendJson(c, 200, "success", nil)
	}
}
