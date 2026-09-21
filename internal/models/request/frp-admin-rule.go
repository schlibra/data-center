package request

type CreateFrpAdminRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
	User  int `json:"user"`
}
type UpdateFrpAdminRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
	User  int `json:"user"`
}
