package openvpn

import (
	"crypto/tls"
	"data-center/pkg/response"
	"data-center/pkg/utils"
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type H map[string]any

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

func closeDB(db *sql.DB) {
	_ = db.Close()
}
func iKuaiRequest(apiUrl string, res any) (*resty.Request, string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return nil, "", err
	}
	token := cfg.IKuai.Key
	url := strings.TrimRight(cfg.IKuai.Address, "/") + "/api/v4.0" + apiUrl
	return resty.
		New().
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).
		R().
		SetResult(&res).
		SetAuthScheme("Bearer").
		SetAuthToken(token), url, nil
}
