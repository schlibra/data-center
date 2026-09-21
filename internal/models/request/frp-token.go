package request

type FrpTokenCreate struct {
	Name string `json:"name"`
}
type FrpTokenInfo struct {
	ID int `uri:"id"`
}
type FrpTokenUpdate struct {
	ID     int    `uri:"id"`
	Name   string `json:"name"`
	Enable int    `json:"enable"`
}
type FrpTokenDelete struct {
	ID int `uri:"id"`
}
type FrpTokenGenerate struct {
	ID int `uri:"id"`
}
