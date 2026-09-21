package request

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

type FrpApiLogin struct {
	Version string             `json:"version"`
	Op      string             `json:"op"`
	Content frpApiLoginContent `json:"content"`
}
type frpApiProxyContentUser struct {
	User string `json:"user"`
}
type frpApiProxyContent struct {
	User       frpApiProxyContentUser `json:"user"`
	RemotePort int                    `json:"remote_port"`
}
type FrpApiProxy struct {
	Content frpApiProxyContent `json:"content"`
}
type FrpApiClientsDataItem struct {
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
	Items    []FrpApiClientsDataItem `json:"items"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Total    int                     `json:"total"`
}
type FrpApiClients struct {
	Code int               `json:"code"`
	Data frpApiClientsData `json:"data"`
	Msg  string            `json:"msg"`
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
type FrpApiProxiesDataItem struct {
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
	Items    []FrpApiProxiesDataItem `json:"items"`
}
type FrpApiProxies struct {
	Code int               `json:"code"`
	Data frpApiProxiesData `json:"data"`
	Msg  string            `json:"msg"`
}
