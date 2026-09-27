package request

type openVPNGroupValue struct {
	IP      string `json:"ip"`
	Comment string `json:"comment"`
}
type CreateOpenVPNGroup struct {
	Name  string              `json:"name"`
	Value []openVPNGroupValue `json:"value"`
}
type UpdateOpenVPNGroup struct {
	Name  string              `json:"name"`
	Value []openVPNGroupValue `json:"value"`
}
