package request

type CreateFrpRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
}
type UpdateFrpRule struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Token int `json:"token"`
}
