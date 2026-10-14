package mail

import (
	"bytes"
	"data-center/internal/models/response"
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/mail"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jhillyerd/enmime"
	"github.com/knadh/go-pop3"
)

func ListMailMessageService(c *gin.Context, domain int, record string) {
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
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		Get(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, record))
	if err != nil {
		sendError(c, err)
		return
	}
	if res.StatusCode() == 200 {
		var result response.MailListRecord
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		p3 := pop3.New(pop3.Opt{
			Host: domainRow.Server,
			Port: domainRow.Port,
			TLSEnabled: func() bool {
				return domainRow.SSL == "1" || domainRow.SSL == "Y"
			}(),
		})
		client, err := p3.NewConn()
		if err != nil {
			sendError(c, err)
		}
		defer func(client *pop3.Conn) {
			_ = client.Quit()
		}(client)
		if err := client.Auth(result.Fields.Email, result.Fields.Password); err != nil {
			sendError(c, err)
		}
		count, size, _ := client.Stat()
		var mails []H
		for id := 1; id <= count; id++ {
			rawMsg, err := client.Retr(id)
			if err != nil {
				sendError(c, err)
			}
			bodyBytes, err := io.ReadAll(rawMsg.Body)
			if err != nil {
				sendError(c, err)
			}
			re := regexp.MustCompile(`([^\r\n])(Content-[a-zA-Z0-9-]+:)`)
			cleanBodyBytes := re.ReplaceAll(bodyBytes, []byte("$1\r\n$2"))
			var fullMailBuf bytes.Buffer
			fullMailBuf.WriteString(fmt.Sprintf("Content-Type: %s\r\n", rawMsg.Header.Get("Content-Type")))
			if cte := rawMsg.Header.Get("Content-Transfer-Encoding"); cte != "" {
				fullMailBuf.WriteString(fmt.Sprintf("Content-Transfer-Encoding: %s\r\n", cte))
			}
			fullMailBuf.WriteString("MIME-Version: 1.0\r\n")
			fullMailBuf.WriteString("\r\n")
			fullMailBuf.Write(cleanBodyBytes)
			env, err := enmime.ReadEnvelope(&fullMailBuf)
			if err != nil {
				sendError(c, err)
			}
			from, err := mail.ParseAddress(rawMsg.Header.Get("From"))
			if err != nil {
				sendError(c, err)
			}
			dec := new(mime.WordDecoder)
			subject, err := dec.DecodeHeader(rawMsg.Header.Get("Subject"))
			if err != nil {
				sendError(c, err)
			}
			mails = append(mails, H{
				"id":      id,
				"subject": subject,
				"from": H{
					"address": from.Address,
					"name":    from.Name,
				},
				"body": H{
					"text": env.Text,
					"html": env.HTML,
				},
			})
		}
		sendJson(c, 200, "Success", H{
			"count": count,
			"size":  size,
			"mails": mails,
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
func GetMailMessageService(c *gin.Context, domain int, record string, id int) {
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
	res, err := resty.
		New().
		R().
		SetAuthScheme("Bearer").
		SetAuthToken(teableKey).
		Get(fmt.Sprintf("%s/api/table/%s/record/%s", teableUrl, domainRow.TableId, record))
	if err != nil {
		sendError(c, err)
		return
	}
	if res.StatusCode() == 200 {
		var result response.MailListRecord
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			sendError(c, err)
		}
		p3 := pop3.New(pop3.Opt{
			Host: domainRow.Server,
			Port: domainRow.Port,
			TLSEnabled: func() bool {
				return domainRow.SSL == "1" || domainRow.SSL == "Y"
			}(),
		})
		client, err := p3.NewConn()
		if err != nil {
			sendError(c, err)
		}
		defer func(client *pop3.Conn) {
			_ = client.Quit()
		}(client)
		if err := client.Auth(result.Fields.Email, result.Fields.Password); err != nil {
			sendError(c, err)
		}
		rawMsg, err := client.Retr(id)
		if err != nil {
			sendError(c, err)
		}
		bodyBytes, err := io.ReadAll(rawMsg.Body)
		if err != nil {
			sendError(c, err)
		}
		re := regexp.MustCompile(`([^\r\n])(Content-[a-zA-Z0-9-]+:)`)
		cleanBodyBytes := re.ReplaceAll(bodyBytes, []byte("$1\r\n$2"))
		var fullMailBuf bytes.Buffer
		fullMailBuf.WriteString(fmt.Sprintf("Content-Type: %s\r\n", rawMsg.Header.Get("Content-Type")))
		if cte := rawMsg.Header.Get("Content-Transfer-Encoding"); cte != "" {
			fullMailBuf.WriteString(fmt.Sprintf("Content-Transfer-Encoding: %s\r\n", cte))
		}
		fullMailBuf.WriteString("MIME-Version: 1.0\r\n")
		fullMailBuf.WriteString("\r\n")
		fullMailBuf.Write(cleanBodyBytes)
		env, err := enmime.ReadEnvelope(&fullMailBuf)
		if err != nil {
			sendError(c, err)
		}
		from, err := mail.ParseAddress(rawMsg.Header.Get("From"))
		if err != nil {
			sendError(c, err)
		}
		dec := new(mime.WordDecoder)
		subject, err := dec.DecodeHeader(rawMsg.Header.Get("Subject"))
		if err != nil {
			sendError(c, err)
		}

		sendJson(c, 200, "Success", H{
			"id":      id,
			"subject": subject,
			"from": H{
				"address": from.Address,
				"name":    from.Name,
			},
			"body": H{
				"text": env.Text,
				"html": env.HTML,
			},
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
