package request

type CreateAdminUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Group    int    `json:"group"`
	Enable   int    `json:"enable"`
}
type UpdateAdminUser struct {
	Nickname string `json:"nickname"`
	Group    int    `json:"group"`
	Enable   int    `json:"enable"`
}
type PasswordAdminUser struct {
	Password string `json:"password"`
}
