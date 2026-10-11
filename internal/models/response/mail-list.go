package response

type MailListRecord struct {
	AuthNumber  int    `json:"authNumber"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime string `json:"createdTime"`
	Fields      struct {
		ID       int    `json:"ID"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	} `json:"fields"`
	ID               string `json:"id"`
	LastModifiedBy   string `json:"lastModifiedBy"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Name             string `json:"name"`
}

type MailList struct {
	Extra struct {
		NextCursor string `json:"nextCursor"`
	} `json:"extra"`
	Records []MailListRecord `json:"records"`
}
type MailListResponse struct {
	Records []MailListRecord `json:"records"`
	Message string           `json:"message"`
	Status  int              `json:"status"`
	Code    string           `json:"code"`
}
