package service

import (
	"data-center/internal/models/request"
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
	var proxyResult request.FrpApiProxies
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&proxyResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/proxies?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", proxyResult.Data.Items)
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
	var clientResult request.FrpApiClients
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&clientResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/clients?pageSize=200", frpsHost, frpsPort))
	sendJson(c, 200, "success", clientResult.Data.Items)
}
