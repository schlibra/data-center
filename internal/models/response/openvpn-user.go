package response

type openVPNUserResultsDataSrcAddrObject struct {
	Type   int    `json:"type"`
	GPName string `json:"gp_name"`
	GID    string `json:"gid"`
}
type openVPNUserResultsDataSrcAddr struct {
	Object []openVPNUserResultsDataSrcAddrObject `json:"object"`
}
type openVPNUserResultsData struct {
	LastOfftime   int                           `json:"last_offtime"`
	CardID        string                        `json:"cardid"`
	ID            int                           `json:"id"`
	Enabled        string                        `json:"enabled"`
	Comment       string                        `json:"comment"`
	Username      string                        `json:"username"`
	Duration      int                           `json:"duration"`
	TagName       string                        `json:"tagname"`
	Expires       int                           `json:"expires"`
	StartTime     int                           `json:"start_time"`
	CreateTime    int                           `json:"create_time"`
	PPPType       string                        `json:"ppptype"`
	Passwd        string                        `json:"passwd"`
	PPPName       string                        `json:"pppname"`
	Share         int                           `json:"share"`
	AutoMac       int                           `json:"auto_mac"`
	Upload        int                           `json:"upload"`
	Download      int                           `json:"download"`
	IPType        int                           `json:"ip_type"`
	SrcAddr       openVPNUserResultsDataSrcAddr `json:"src_addr"`
	Mac           string                        `json:"mac"`
	Address       string                        `json:"address"`
	Name          string                        `json:"name"`
	Phone         string                        `json:"phone"`
	Packages      int                           `json:"packages"`
	ProxyUsername string                        `json:"proxy_username"`
	PPPoEV6Wan    string                        `json:"pppoev6_wan"`
	AutoVlanID    int                           `json:"auto_vlanid"`
	BindVlanID    string                        `json:"bind_vlanid"`
	BindIfName    string                        `json:"bind_ifname"`
	LastConnTime  int                           `json:"last_conn_time"`
}
type openVPNUserResults struct {
	Data  []openVPNUserResultsData `json:"data"`
	Total int                      `json:"total"`
}
type OpenVPNUser struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Results openVPNUserResults `json:"results"`
}
type OpenVPNUserChange struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	RowID   int    `json:"row_id"`
}
