package frp_admin

type CreateFrpAdminToken struct {
	Name string `json:"name"`
	User int    `json:"user"`
}
type UpdateFrpAdminToken struct {
	Name   string `json:"name"`
	User   int    `json:"user"`
	Enable int    `json:"enable"`
}
