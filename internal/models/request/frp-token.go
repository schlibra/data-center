package request

type FrpTokenCreate struct {
	Name string `json:"name"`
}
type FrpTokenUpdate struct {
	Name   string `json:"name"`
	Enable int    `json:"enable"`
}
