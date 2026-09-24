package request

type CreateAdminGroup struct {
	Name       string `json:"name"`
	Admin      int    `json:"admin"`
	Permission string `json:"permission"`
}
type UpdateAdminGroup struct {
	Name       string `json:"name"`
	Admin      int    `json:"admin"`
	Permission string `json:"permission"`
}
