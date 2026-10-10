package mail

import (
	"data-center/internal/models/request/mail"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListMailDomainService(c *gin.Context) {
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	domains, err := mailDomain.SelectAll()
	sendI18n(c, 200, "mail.domain.get_success", domains)
}
func GetMailDomainService(c *gin.Context, id int) {
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	domain, err := mailDomain.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "mail.domain.not_exist", nil)
	}
	sendI18n(c, 200, "mail.domain.get_success", domain)
}
func CreateMailDomainService(c *gin.Context, req mail.CreateMailDomain) {
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	if _, err := mailDomain.SelectByDomain(req.Domain); err == nil {
		sendI18n(c, 400, "mail.domain.exist", nil)
	}
	if nw(c) {
		_, err := mailDomain.Insert(req.Domain, req.Name, req.TableId)
		if err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "mail.domain.create_success", nil)
	}
}
func UpdateMailDomainService(c *gin.Context, id int, req mail.UpdateMailDomain) {
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	row, err := mailDomain.SelectById(id)
	if err != nil {
		sendI18n(c, 400, "mail.domain.not_exist", nil)
		return
	}
	if nw(c) {
		if _, err := mailDomain.UpdateById(id, row.Domain, req.Name, req.TableId); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "mail.domain.update_success", nil)
	}
}
func DeleteMailDomainService(c *gin.Context, id int) {
	mailDomain, err := repository.NewMailDomain()
	if err != nil {
		sendError(c, err)
	}
	defer closeDB(mailDomain.DB)
	if _, err := mailDomain.SelectById(id); err != nil {
		sendI18n(c, 400, "mail.domain.not_exist", nil)
	}
	if nw(c) {
		if _, err := mailDomain.DeleteById(id); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "mail.domain.delete_success", nil)
	}
}
