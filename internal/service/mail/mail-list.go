package mail

import (
	"data-center/internal/models/request/mail"
	"data-center/internal/models/response"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"encoding/json"
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
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
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
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		Get(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, id))
	if err != nil {
		sendError(c, err)
	}
	if res.StatusCode() == 200 {
		var result response.MailListRecord
		err = json.Unmarshal(res.Body(), &result)
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
	} else {
		var result response.MailListResponse
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, result.Status, result.Message, result)
	}
}
func CreateMailListService(c *gin.Context, domain int, req mail.CreateMailList) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
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
	if res.StatusCode() == 201 {
		var result response.MailList
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, 200, "Success", result)
	} else {
		var result response.MailListResponse
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, result.Status, result.Message, result)
	}
}
func UpdateMailListService(c *gin.Context, domain int, id string, req mail.UpdateMailList) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		SetBody(H{
			"record": H{
				"fields": req,
			},
		}).Patch(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, id))
	if err != nil {
		sendError(c, err)
	}
	if res.StatusCode() == 200 {
		var result response.MailListUpdateResponse
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, 200, "Success", result)
	} else {
		var result response.MailListResponse
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, result.Status, result.Message, result)
	}

}
func DeleteMailListService(c *gin.Context, domain int, id string) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		sendError(c, err)
	}
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	domainRow, err := mailDomain.SelectById(domain)
	if err != nil {
		sendError(c, err)
	}
	teableUrl := strings.TrimRight(cfg.Teable.Address, "/")
	teableKey := cfg.Teable.Key
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		Delete(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, id))
	if err != nil {
		sendError(c, err)
	}
	if res.StatusCode() == 200 {
		sendJson(c, 200, "Success", nil)
	} else {
		var result response.MailListResponse
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		sendJson(c, result.Status, result.Message, result)
	}
}
