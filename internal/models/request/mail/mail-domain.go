package mail

type CreateMailDomain struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	TableId string `json:"tableId"`
}
type UpdateMailDomain struct {
	Name    string `json:"name"`
	TableId string `json:"tableId"`
}
