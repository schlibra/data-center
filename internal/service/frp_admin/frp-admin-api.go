package frp_admin

import (
	"data-center/internal/models/request/frp"
	"data-center/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func ProxiesFrpAdminApiService(c *gin.Context) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	client := resty.New()
	var proxyResult frp.ProxiesFrpApi
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&proxyResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/proxies?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", proxyResult.Data.Items)
}

func ClearOfflineFrpAdminApiService(c *gin.Context) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpToken := cfg.FrpsWeb.Token
	_, err = resty.New().R().SetAuthScheme("Basic").SetAuthScheme(frpToken).Post(fmt.Sprintf("http://%s:%d/api/v2/system/prune?type=offline", frpsHost, frpsPort))
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "success", nil)
}

func ClientsFrpAdminApiService(c *gin.Context) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	client := resty.New()
	var clientResult frp.ClientsFrpApi
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&clientResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/clients?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", clientResult.Data.Items)
}
