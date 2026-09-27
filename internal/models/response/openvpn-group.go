package response

type openVPNGroupResultIpDataGroupValue struct {
	IP string `json:"ip"`
}
type openVPNGroupResultIpData struct {
	RefCount   int                                  `json:"ref_count"`
	ID         int                                  `json:"id"`
	GroupName  string                               `json:"group_name"`
	GroupValue []openVPNGroupResultIpDataGroupValue `json:"group_value"`
}
type OpenVPNGroupResult struct {
	IpData  []openVPNGroupResultIpData `json:"ip_data"`
	IpTotal int                        `json:"ip_total"`
}
type OpenVPNGroupChange struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	RowID   int    `json:"rowid"`
}
type OpenVPNGroup struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Results OpenVPNGroupResult `json:"results"`
}
