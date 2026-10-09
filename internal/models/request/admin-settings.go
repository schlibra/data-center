package request

type CreateAdminSettings struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type UpdateAdminSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
