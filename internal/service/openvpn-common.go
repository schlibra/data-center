package service

import (
	"crypto/tls"
	"data-center/pkg/utils"
	"strings"

	"github.com/go-resty/resty/v2"
)

func iKuaiRequest(apiUrl string, res any) (*resty.Request, string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return nil, "", err
	}
	token := cfg.IKuai.Key
	url := strings.TrimRight(cfg.IKuai.Address, "/") + "/api/v4.0" + apiUrl
	return resty.
		New().
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).
		R().
		SetResult(&res).
		SetAuthScheme("Bearer").
		SetAuthToken(token), url, nil
}
