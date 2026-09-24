package service

import (
	"data-center/internal/models"
	"data-center/internal/models/request"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func LoginFrpApiService(c *gin.Context, req request.FrpApiLogin) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		c.JSON(200, H{
			"reject":        true,
			"reject_reason": err.Error(),
		})
		return
	}
	defer closeDB(frpToken.DB)
	user := req.Content.User
	token := req.Content.Metas.Token
	tokenRow, err := frpToken.SelectByName(user)
	if err != nil {
		c.JSON(200, H{
			"reject":        true,
			"reject_reason": "token not exist",
		})
		return
	}
	if tokenRow.Token == token {
		c.JSON(200, H{
			"reject":   false,
			"unchange": true,
		})
		return
	}
	if tokenRow.Enable == 0 {
		c.JSON(200, H{
			"reject":        true,
			"reject_reason": "token not enable",
		})
	}
	c.JSON(200, H{
		"reject":        true,
		"reject_reason": "token not match",
	})
}
func ProxyFrpApiService(c *gin.Context, req request.FrpApiProxy) {
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		c.JSON(200, H{
			"reject":        true,
			"reject_reason": err.Error(),
		})
		return
	}
	defer closeDB(frpToken.DB)
	frpRule, err := repository.NewFrpRule()
	if err != nil {
		c.JSON(200, gin.H{
			"reject":        true,
			"reject_reason": err.Error(),
		})
		return
	}
	user := req.Content.User.User
	tokenRow, err := frpToken.SelectByName(user)
	if err != nil {
		c.JSON(200, H{
			"reject":        true,
			"reject_reason": "token not exist",
		})
		return
	}
	ruleRows, err := frpRule.SelectByToken(tokenRow.ID)
	allowPort := make([]string, 0)
	remotePort := req.Content.RemotePort
	for _, ruleRow := range ruleRows {
		if remotePort >= ruleRow.Min && remotePort <= ruleRow.Max {
			c.JSON(200, H{
				"reject":   false,
				"unchange": true,
			})
		}
		allowPort = append(allowPort, strconv.Itoa(ruleRow.Min)+func() string {
			if ruleRow.Max == ruleRow.Min {
				return ""
			}
			return "-" + strconv.Itoa(ruleRow.Max)
		}())
	}
	ports := strings.Join(allowPort, ",")
	c.JSON(200, H{
		"reject":        true,
		"reject_reason": "port not allow, allow ports: " + ports,
	})

}
func ProxiesFrpApiService(c *gin.Context, row models.UserTable) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	tokens, err := frpToken.SelectByUser(row.ID)
	var tokenList []string
	for _, token := range tokens {
		tokenList = append(tokenList, token.Name)
	}
	client := resty.New()
	var proxyResult request.FrpApiProxies
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&proxyResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/proxies?pageSize=200", frpsHost, frpsPort))
	if err != nil {
		sendError(c, err)
	}
	if proxyResult.Code != 200 {
		sendJson(c, proxyResult.Code, proxyResult.Msg, proxyResult.Data)
	}
	proxies := make([]request.FrpApiProxiesDataItem, 0)
	for _, item := range proxyResult.Data.Items {
		if slices.Contains(tokenList, item.User) {
			proxies = append(proxies, item)
		}
	}
	sendJson(c, 200, "success", proxies)
}
func ClientFrpApiService(c *gin.Context, row models.UserTable) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	frpsHost := cfg.FrpsWeb.Host
	frpsPort := cfg.FrpsWeb.Port
	frpsToken := cfg.FrpsWeb.Token
	frpToken, err := repository.NewFrpToken()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(frpToken.DB)
	tokens, err := frpToken.SelectByUser(row.ID)
	var tokenList []string
	for _, token := range tokens {
		tokenList = append(tokenList, token.Name)
	}
	client := resty.New()
	var clientResult request.FrpApiClients
	_, err = client.R().
		SetAuthScheme("Basic").
		SetAuthToken(frpsToken).
		SetResult(&clientResult).
		Get(fmt.Sprintf("http://%s:%d/api/v2/clients?pageSize=200", frpsHost, frpsPort))
	if err != nil {
		sendError(c, err)
	}
	if clientResult.Code != 200 {
		sendJson(c, clientResult.Code, clientResult.Msg, clientResult.Data)
	}
	clients := make([]request.FrpApiClientsDataItem, 0)
	for _, item := range clientResult.Data.Items {
		if slices.Contains(tokenList, item.User) {
			clients = append(clients, item)
		}
	}
	sendJson(c, 200, "success", clients)
}
