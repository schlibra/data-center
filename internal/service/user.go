package service

import (
	"data-center/internal/models"
	"data-center/internal/models/request"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"encoding/base64"
	"encoding/json"
	"math"
	"slices"
	"time"

	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

func UserLoginService(c *gin.Context, req request.UserLogin) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	redis, err := utils.NewRedis()
	if err != nil {
		sendError(c, err)
	}
	row, err := user.SelectByUsername(req.Username)
	if err != nil {
		sendI18n(c, 400, "user.login.not_exists", H{
			"error": err.Error(),
		})
	}
	decodeStr, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		sendError(c, err)
	}
	priKey, err := redis.Get("login_" + req.Username + "_private_key")
	if err != nil {
		sendError(c, err)
	}
	if _, err := redis.Del("login_" + req.Username + "_private_key"); err != nil {
		sendError(c, err)
	}
	priKeyPem := priKey.(string)
	data, err := utils.DecryptWithPrivateKey(decodeStr, priKeyPem)
	if err != nil {
		sendError(c, err)
	}
	if utils.PasswordCheck(row.Password, string(data)) {
		if row.Enable == 1 {
			tokenId, err := utils.GenerateRandomString(32)
			if err != nil {
				sendError(c, err)
			}
			token, err := utils.JwtUserGenerate(row.ID, row.Username, row.Group, row.Enable, tokenId, 24*time.Hour)
			if err != nil {
				sendError(c, err)
			}
			_, err = user.UpdateById(row.ID, row.Password, row.Nickname, row.Group, row.Enable, tokenId, row.ApiId)
			if err != nil {
				sendError(c, err)
			}
			sendJson(c, 200, i18n.MustGetMessage(c, "user.login.success"), H{
				"token": token,
			})
		} else {
			sendI18n(c, 400, "user.login.disabled", nil)
		}
	} else {
		sendI18n(c, 400, "user.login.password_error", nil)
	}

}
func UserLoginKeyService(c *gin.Context, req request.UserLoginKey) {
	redis, err := utils.NewRedis()
	if err != nil {
		sendError(c, err)
	}
	priKey, pubKey, err := utils.GenerateRsaKeyPair()
	if err != nil {
		sendError(c, err)
	}
	err = redis.Set("login_"+req.Username+"_private_key", priKey)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, i18n.MustGetMessage(c, "user.rsa_key.generate_pubkey_success"), H{
		"public_key": pubKey,
	})
}
func UserRegisterService(c *gin.Context, req request.UserRegister) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	defaultGroup := cfg.Server.DefaultGroup
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	if _, err := user.SelectByUsername(req.Username); err == nil {
		sendJson(c, 400, i18n.MustGetMessage(c, "user.register.exists"), nil)
	}
	redis, err := utils.NewRedis()
	if err != nil {
		sendError(c, err)
	}
	decodeStr, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		sendError(c, err)
	}
	priKey, err := redis.Get("register_" + req.Username + "_private_key")
	if err != nil {
		sendError(c, err)
	}
	priKeyPem := priKey.(string)
	data, err := utils.DecryptWithPrivateKey(decodeStr, priKeyPem)
	if err != nil {
		sendError(c, err)
	}
	password, err := utils.PasswordHash(string(data))
	if err != nil {
		sendError(c, err)
	}
	if !c.Writer.Written() {
		_, err = user.Insert(req.Username, password, req.Nickname, defaultGroup, 0)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, 200, i18n.MustGetMessage(c, "user.register.success"), nil)
	}
}
func UserRegisterKeyService(c *gin.Context, req request.UserRegisterKey) {
	redis, err := utils.NewRedis()
	if err != nil {
		sendError(c, err)
	}
	priKey, pubKey, err := utils.GenerateRsaKeyPair()
	if err != nil {
		sendError(c, err)
	}
	err = redis.Set("register_"+req.Username+"_private_key", priKey)
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, i18n.MustGetMessage(c, "user.rsa_key.generate_pubkey_success"), H{
		"public_key": pubKey,
	})
}
func UserLogoutService(c *gin.Context, row models.UserTable) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	_, err = user.UpdateById(row.ID, row.Password, row.Nickname, row.Group, row.Enable, "", row.ApiId)
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "user.logout.success", nil)
}
func UserInfoService(c *gin.Context, row models.UserTable) {
	permission, err := repository.NewPermission()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(permission.DB)
	if row.Enable == 1 {
		var groupPermission []int
		if err := json.Unmarshal([]byte(row.GroupInfo.Permission), &groupPermission); err != nil {
			sendError(c, err)
		}
		permissionList := make([]map[string]string, 0)
		permissions, err := permission.SelectAll()
		if err != nil {
			sendError(c, err)
		}
		for _, p := range permissions {
			if slices.Contains(groupPermission, p.ID) {
				permissionList = append(permissionList, map[string]string{
					"key":  p.Key,
					"name": p.Name,
				})
			}
		}
		row.Password = "********"
		row.TokenID = "********"
		row.ApiId = "********"
		row.Permissions = permissionList
		sendJson(c, 200, i18n.MustGetMessage(c, "user.info.success"), row)
	} else {
		sendI18n(c, 403, "user.login.disabled", nil)
	}
}
func UserUpdateService(c *gin.Context, req request.UserUpdate, row models.UserTable) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	if !c.Writer.Written() {
		_, err = user.UpdateById(row.ID, row.Password, req.Nickname, row.Group, row.Enable, row.TokenID, row.ApiId)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "user.update.success", nil)
	}
}
func UserPasswordService(c *gin.Context, req request.UserPassword, row models.UserTable) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	pwd, err := utils.PasswordHash(req.Password)
	if err != nil {
		sendError(c, err)
	}
	if !c.Writer.Written() {
		_, err = user.UpdateById(row.ID, pwd, row.Nickname, row.Group, row.Enable, "", row.ApiId)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "user.password.success", nil)
	}
}
func GenerateUserApiKeyService(c *gin.Context, row models.UserTable) {
	user, err := repository.NewUser()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(user.DB)
	apiId, err := utils.GenerateRandomString(32)
	if err != nil {
		sendError(c, err)
	}
	token, err := utils.JwtUserGenerate(row.ID, row.Username, row.Group, row.Enable, apiId, time.Duration(math.MaxInt64))
	if err != nil {
		sendError(c, err)
	}
	if !c.Writer.Written() {
		if _, err := user.UpdateById(row.ID, row.Password, row.Nickname, row.Group, row.Enable, row.TokenID, apiId); err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "user.apikey.success", H{
				"token": token,
			})
		}
	}
}

func GetUserApiKeyService(c *gin.Context, row models.UserTable) {
	if row.ApiId == "" {
		sendI18n(c, 400, "user.apikey.not_exists", nil)
	}
	token, err := utils.JwtUserGenerate(row.ID, row.Username, row.Group, row.Enable, row.ApiId, time.Duration(math.MaxInt64))
	if err != nil {
		sendError(c, err)
	}
	sendI18n(c, 200, "user.apikey.success", H{
		"token": token,
	})
}
