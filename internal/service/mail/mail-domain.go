package mail

import (
	"data-center/internal/models/request/mail"
	"data-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListMailDomainService(c *gin.Context) {
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	if domains, err := mailDomain.SelectAll(); err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "mail.domain.get_success", domains)
	}
}
func GetMailDomainService(c *gin.Context, id int) {
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	if domain, err := mailDomain.SelectById(id); err != nil {
		sendI18n(c, 400, "mail.domain.not_exist", nil)
	} else {
		sendI18n(c, 200, "mail.domain.get_success", domain)
	}
}
func CreateMailDomainService(c *gin.Context, req mail.CreateMailDomain) {
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	if nw(c) {
		if _, err := mailDomain.Insert(req.Domain, req.Name, req.TableId, req.Server, req.Port, req.SSL); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "mail.domain.create_success", nil)
	}
}
func UpdateMailDomainService(c *gin.Context, id int, req mail.UpdateMailDomain) {
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
	if _, err := mailDomain.SelectById(id); err != nil {
		sendI18n(c, 400, "mail.domain.not_exist", nil)
	}
	if nw(c) {
		if _, err := mailDomain.UpdateById(id, req.Domain, req.Name, req.TableId, req.Server, req.Port, req.SSL); err != nil {
			sendError(c, err)
		}
		sendI18n(c, 200, "mail.domain.update_success", nil)
	}
}
func DeleteMailDomainService(c *gin.Context, id int) {
	mailDomain := new(repository.MailDomain)
	if err := mailDomain.New(); err != nil {
		sendError(c, err)
	}
	defer mailDomain.Close()
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
