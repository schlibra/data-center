package request

type createOpenVPNUserSrcAddrObject struct {
	Type   int    `json:"type"`
	GPName string `json:"gp_name"`
	GID    string `json:"gid"`
}
type createOpenVPNUserSrcAddr struct {
	Object []createOpenVPNUserSrcAddrObject `json:"object"`
}

type CreateOpenVPNUser struct {
	Username string                   `json:"username"`
	Password string                   `json:"password"`
	Enabled  string                   `json:"enabled"`
	Share    int                      `json:"share"`
	SrcAddr  createOpenVPNUserSrcAddr `json:"src_addr"`
}
type UpdateOpenVPNUser struct {
	Username string                   `json:"username"`
	Password string                   `json:"password"`
	Enabled  string                   `json:"enabled"`
	Share    int                      `json:"share"`
	SrcAddr  createOpenVPNUserSrcAddr `json:"src_addr"`
}
