package request

type CreateAdminPermission struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}
type UpdateAdminPermission struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}
