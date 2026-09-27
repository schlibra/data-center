package service

import (
	"data-center/internal/models/request"
	"data-center/internal/models/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListOpenVPNUserService(c *gin.Context) {
	var result response.OpenVPNUser
	iKuai, url, err := iKuaiRequest("/auth/users", &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.user.get_success", result)
	}
}

func GetOpenVPNUserService(c *gin.Context, id int) {
	var result response.OpenVPNUser
	iKuai, url, err := iKuaiRequest("/auth/users/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.user.get_success", result)
	}
}

func CreateOpenVPNUserService(c *gin.Context, req request.CreateOpenVPNUser) {
	var result response.OpenVPNUserChange
	iKuai, url, err := iKuaiRequest("/auth/users", &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err = iKuai.
			SetBody(H{
				"username":    req.Username,
				"passwd":      req.Password,
				"enabled":     req.Enabled,
				"expires":     0,
				"start_time":  0,
				"ppptype":     "any",
				"share":       req.Share,
				"auto_mac":    1,
				"upload":      0,
				"download":    0,
				"packages":    0,
				"auto_vlanid": 1,
				"bind_vlanid": "0",
				"bind_ifname": "any",
				"ip_type":     1,
				"src_addr":    req.SrcAddr,
			}).
			Post(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.user.create_success", result)
		}
	}
}

func UpdateOpenVPNUserService(c *gin.Context, id int, req request.UpdateOpenVPNUser) {
	var result response.OpenVPNUserChange
	iKuai, url, err := iKuaiRequest("/auth/users/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err = iKuai.SetBody(H{
			"username":    req.Username,
			"passwd":      req.Password,
			"enabled":     req.Enabled,
			"expires":     0,
			"start_time":  0,
			"ppptype":     "any",
			"share":       req.Share,
			"auto_mac":    1,
			"upload":      0,
			"download":    0,
			"packages":    0,
			"auto_vlanid": 1,
			"bind_vlanid": "0",
			"bind_ifname": "any",
			"ip_type":     1,
			"comment":     "",
			"pppname":     "",
			"mac":         "",
			"address":     "",
			"name":        "",
			"phone":       "",
			"cardid":      "",
			"pppoev6_wan": "",
			"src_addr":    req.SrcAddr,
		}).Put(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.user.update_success", result)
		}
	}
}

func DeleteOpenVPNUserService(c *gin.Context, id int) {
	var result response.OpenVPNUserChange
	iKuai, url, err := iKuaiRequest("/auth/users/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err = iKuai.Delete(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.user.delete_success", result)
		}
	}
}
