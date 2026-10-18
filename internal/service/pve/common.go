package pve

import (
	"crypto/tls"
	"data-center/pkg/response"
	"data-center/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type H gin.H

func sendJson(c *gin.Context, code int, message string, data any) {
	response.SendJson(c, code, message, data)
}
func sendI18n(c *gin.Context, code int, messageId string, data any) {
	response.SendI18n(c, code, messageId, data)
}
func sendError(c *gin.Context, err error) {
	response.SendError(c, err)
}
func nw(c *gin.Context) bool {
	return response.NW(c)
}
func pveRequest(apiUrl string, res any) (*resty.Request, string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return nil, "", err
	}
	url := strings.TrimLeft(cfg.Pve.Address, "/") + "/api2/json" + apiUrl
	return resty.
		New().
		SetDebug(cfg.Server.Debug).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		R().
		SetResult(&res).
		SetHeader("Authorization", "PVEAPIToken="+cfg.Pve.Id+"="+cfg.Pve.Token), url, nil
}
