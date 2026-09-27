package response

type openVPNClientResultsData struct {
	Session      string `json:"session"`
	UID          string `json:"uid"`
	PPPDev       string `json:"pppdev"`
	Mac          string `json:"mac"`
	Expires      int    `json:"expires"`
	ID           int    `json:"id"`
	PackName     string `json:"packname"`
	Upload       int    `json:"upload"`
	CheckVlanRes int    `json:"check_vlan_res"`
	IpAddrInt    string `json:"ip_addr_int"`
	Download     int    `json:"download"`
	Name         string `json:"name"`
	WebID        int    `json:"webid"`
	Phone        string `json:"phone"`
	Comment      string `json:"comment"`
	Packages     int    `json:"packages"`
	Interface    string `json:"interface"`
	IpAddr       string `json:"ip_addr"`
	Username     string `json:"username"`
	PPPType      string `json:"ppptype"`
	AuthTime     int    `json:"auth_time"`
}
type openVPNClientResults struct {
	Data  []openVPNClientResultsData `json:"data"`
	Total int                        `json:"total"`
}
type OpenVPNClient struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Results openVPNClientResults `json:"results"`
}
type OpenVPNClientKick struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
