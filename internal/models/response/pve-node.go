package response

type PveNode struct {
	Data []struct {
		Id             string  `json:"id"`
		Mem            int     `json:"mem"`
		Cpu            float64 `json:"cpu"`
		Level          string  `json:"level"`
		MaxMem         int     `json:"maxmem"`
		Disk           int     `json:"disk"`
		Node           string  `json:"node"`
		Type           string  `json:"type"`
		SSLFingerPrint string  `json:"ssl_fingerprint"`
		MaxCpu         int     `json:"maxcpu"`
		Uptime         int     `json:"uptime"`
		MaxDisk        int     `json:"maxdisk"`
		Status         string  `json:"status"`
	} `json:"data"`
}
