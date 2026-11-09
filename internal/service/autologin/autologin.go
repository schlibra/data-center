package autologin

import (
	"data-center/internal/models"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListAutoLoginService(c *gin.Context) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	rows, err := autoLogin.SelectAll()
	if err != nil {
		sendError(c, err)
	}
	data := make([]models.AutoLoginTable, 0)
	for _, row := range rows {
		row.Username = ""
		row.Password = ""
		data = append(data, row)
	}
	sendJson(c, 200, "Success", data)
}
func GetAutoLoginService(c *gin.Context, id int) {
	autoLogin := new(repository.AutoLogin)
	if err := autoLogin.New(); err != nil {
		sendError(c, err)
	}
	defer autoLogin.Close()
	row, err := autoLogin.SelectById(id)
	if err != nil {
		sendJson(c, 400, "not found", H{
			"error": err.Error(),
		})
	}
	key, err := utils.GenerateRandomString(8)
	if err != nil {
		sendError(c, err)
	}
	keyByte := []byte(key)
	jsonData := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s", "url": "%s", "site": "%s"}`, row.Username, row.Password, row.URL, row.Site))
	enc, err := utils.AesCBCEncrypt(jsonData, keyByte, keyByte)
	if err != nil {
		sendError(c, err)
	}
	encStr := base64.StdEncoding.EncodeToString(enc)
	autologinJson := []byte(fmt.Sprintf(`{"data": "%s", "key": "%s"}`, encStr, key))
	autoLoginStr := strings.TrimRight(base64.StdEncoding.EncodeToString(autologinJson), "=")
	sendJson(c, 200, "Success", H{
		"url":  row.URL,
		"data": autoLoginStr,
	})
}
