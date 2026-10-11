package frp

import (
	"data-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type frpApiLoginContentMetas struct {
	Token string `json:"token"`
}
type frpApiLoginContent struct {
	Arch          string                  `json:"arch"`
	ClientAddress string                  `json:"client_address"`
	ClientSpec    any                     `json:"client_spec"`
	Hostname      string                  `json:"hostname"`
	Metas         frpApiLoginContentMetas `json:"metas"`
	OS            string                  `json:"os"`
	PoolCount     int                     `json:"pool_count"`
	PrivilegeKey  string                  `json:"privilege_key"`
	RunID         string                  `json:"run_id"`
	Timestamp     int                     `json:"timestamp"`
	User          string                  `json:"user"`
	Version       string                  `json:"version"`
}

type LoginFrpApi struct {
	Version string             `json:"version"`
	Op      string             `json:"op"`
	Content frpApiLoginContent `json:"content"`
}

func (lfa LoginFrpApi) Parse(c *gin.Context) LoginFrpApi {
	if err := c.ShouldBindJSON(&lfa); err != nil {
		response.SendError(c, err)
	}
	return lfa
}

type frpApiProxyContentUser struct {
	User string `json:"user"`
}
type frpApiProxyContent struct {
	User       frpApiProxyContentUser `json:"user"`
	RemotePort int                    `json:"remote_port"`
}
type ProxyFrpApi struct {
	Content frpApiProxyContent `json:"content"`
}

func (pfa ProxyFrpApi) Parse(c *gin.Context) ProxyFrpApi {
	if err := c.ShouldBindJSON(&pfa); err != nil {
		response.SendError(c, err)
	}
	return pfa
}

type ClientsFrpApiDataItem struct {
	ClientID         string `json:"clientID"`
	ClientIP         string `json:"clientIP"`
	FirstConnectedAt int64  `json:"firstConnectedAt"`
	Hostname         string `json:"hostname"`
	Key              string `json:"key"`
	LastConnectedAt  int64  `json:"lastConnectedAt"`
	Online           bool   `json:"online"`
	RunID            string `json:"runID"`
	User             string `json:"user"`
	Version          string `json:"version"`
	WireProtocol     string `json:"wireProtocol"`
}
type frpApiClientsData struct {
	Items    []ClientsFrpApiDataItem `json:"items"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Total    int                     `json:"total"`
}
type ClientsFrpApi struct {
	Code int               `json:"code"`
	Data frpApiClientsData `json:"data"`
	Msg  string            `json:"msg"`
}

func (cfa ClientsFrpApi) Parse(c *gin.Context) ClientsFrpApi {
	if err := c.ShouldBindJSON(&cfa); err != nil {
		response.SendError(c, err)
	}
	return cfa
}

type frpApiProxiesDataItemSpecTypeTransport struct {
	UseEncryption      bool   `json:"useEncryption"`
	UseCompression     bool   `json:"useCompression"`
	BandWidthLimit     string `json:"bandWidthLimit"`
	BandWidthLimitMode string `json:"bandWidthLimitMode"`
}
type frpApiProxiesDataItemSpecTypeLoadBalancer struct {
	Group string `json:"group"`
}
type frpApiProxiesDataItemSpecType struct {
	Transport    frpApiProxiesDataItemSpecTypeTransport    `json:"transport"`
	LoadBalancer frpApiProxiesDataItemSpecTypeLoadBalancer `json:"loadBalancer"`
	RemotePort   int                                       `json:"remotePort"`
}
type frpApiProxiesDataItemSpec struct {
	Type string                        `json:"type"`
	Tcp  frpApiProxiesDataItemSpecType `json:"tcp"`
	Udp  frpApiProxiesDataItemSpecType `json:"udp"`
}
type frpApiProxiesDataItemStatus struct {
	Phase           string `json:"phase"`
	TodayTrafficIn  int64  `json:"todayTrafficIn"`
	TodayTrafficOut int64  `json:"todayTrafficOut"`
	CurConns        int    `json:"curConns"`
	LastStartAt     int64  `json:"lastStartAt"`
	LastCloseAt     int64  `json:"lastCloseAt"`
}
type ProxiesFrpApiDataItem struct {
	Name     string                      `json:"name"`
	User     string                      `json:"user"`
	ClientID string                      `json:"clientID"`
	Spec     frpApiProxiesDataItemSpec   `json:"spec"`
	Status   frpApiProxiesDataItemStatus `json:"status"`
}

type frpApiProxiesData struct {
	Total    int                     `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Items    []ProxiesFrpApiDataItem `json:"items"`
}
type ProxiesFrpApi struct {
	Code int               `json:"code"`
	Data frpApiProxiesData `json:"data"`
	Msg  string            `json:"msg"`
}

func (pfa ProxiesFrpApi) Parse(c *gin.Context) ProxiesFrpApi {
	if err := c.ShouldBindJSON(&pfa); err != nil {
		response.SendError(c, err)
	}
	return pfa
}
