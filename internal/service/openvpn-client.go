package service

import (
	"data-center/internal/models/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListOpenVPNClientService(c *gin.Context) {
	var result response.OpenVPNClient
	iKuai, url, err := iKuaiRequest("/auth/online-users", &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.client.get_success", result)
	}
}

func GetOpenVPNClientService(c *gin.Context, id int) {
	var result response.OpenVPNClient
	iKuai, url, err := iKuaiRequest("/auth/online-users/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.client.get_success", result)
	}
}

func KickOpenVPNClientService(c *gin.Context, id int) {
	var result response.OpenVPNClientKick
	iKuai, url, err := iKuaiRequest("/auth/online-users/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err = iKuai.Delete(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.client.kick_success", result)
		}
	}
}
