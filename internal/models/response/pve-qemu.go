package response

type PveQemu struct {
	Data []struct {
		PressureCpuSome    float64 `json:"pressurecpusome"`
		PressureMemorySome float64 `json:"pressurememorysome"`
		PressureIoSome     float64 `json:"pressureiosome"`
		PressureCpuFull    float64 `json:"pressurecpufull"`
		PressureMemoryFull float64 `json:"pressurememoryfull"`
		PressureIoFull     float64 `json:"pressureiofull"`
		Mem                float64 `json:"mem"`
		MaxMem             float64 `json:"maxmem"`
		MemHost            float64 `json:"memhost"`
		Disk               float64 `json:"disk"`
		MaxDisk            float64 `json:"maxdisk"`
		Cpu                float64 `json:"cpu"`
		Cpus               float64 `json:"cpus"`
		NetIn              int     `json:"netin"`
		NetOut             int     `json:"netout"`
		Status             string  `json:"status"`
		Uptime             int     `json:"uptime"`
		Pid                int     `json:"pid"`
		Tags               string  `json:"tags"`
		VmId               int     `json:"vmid"`
		Name               string  `json:"name"`
	} `json:"data"`
}
