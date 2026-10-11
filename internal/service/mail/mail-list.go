package mail

import (
	"data-center/internal/models/request/mail"
	"data-center/internal/models/response"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func ListMailListService(c *gin.Context, domain int) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	var result response.MailList
	_, err = resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		SetResult(&result).
		Get(fmt.Sprintf("%s/api/table/%s/record?take=1000", teableUrl, domainRow.TableId))
	if err != nil {
		sendError(c, err)
	}
	var results []H
	for _, item := range result.Records {
		results = append(results, H{
			"id":       item.ID,
			"ID":       item.Fields.ID,
			"username": item.Fields.Username,
			"password": item.Fields.Password,
			"email":    item.Fields.Email,
		})
	}
	sendJson(c, 200, "Success", results)
}
func GetMailListService(c *gin.Context, domain int, id string) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	var result response.MailListRecord
	_, err = resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		SetResult(&result).
		Get(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, id))
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "Success", H{
		"id":       result.ID,
		"ID":       result.Fields.ID,
		"username": result.Fields.Username,
		"password": result.Fields.Password,
		"email":    result.Fields.Email,
	})
}
func CreateMailListService(c *gin.Context, domain int, req mail.CreateMailList) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	var result any
	_, err = resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		SetResult(&result).
		SetBody(H{
			"records": []H{
				{
					"fields": req,
				},
			},
		}).
		Post(fmt.Sprintf("%s/api/table/%s/record", teableUrl, domainRow.TableId))
	if err != nil {
		sendError(c, err)
	}
	//log.Println(resp.String())
	sendJson(c, 200, "Success", result)
}
func UpdateMailListService(c *gin.Context, domain int, id string, req mail.UpdateMailList) {

}
func DeleteMailListService(c *gin.Context, domain int, id string) {
	//cfg, err := utils.LoadConfig()
	//if err != nil {
	//	sendError(c, err)
	//}
	//mailDomain, err := repository.NewMailDomain()
	//if err != nil {
	//	sendError(c, err)
	//}
	//defer closeDB(mailDomain.DB)
	//domainRow, err := mailDomain.SelectById(domain)
	//if err != nil {
	//	sendError(c, err)
	//}
	//teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	//teableKey := cfg.Teable.Key
	//var
	//_, err = resty.
	//	New().
	//	R().
	//	SetAuthScheme("Bearer").
	//	SetAuthToken(teableKey).
	//	SetResult(&any{}).
}
