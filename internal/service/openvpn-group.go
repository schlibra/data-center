package service

import (
	"data-center/internal/models/request"
	"data-center/internal/models/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListOpenVPNGroupService(c *gin.Context) {
	var result response.OpenVPNGroup
	iKuai, url, err := iKuaiRequest("/ip-objects", &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.
		Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.group.get_success", result.Results)
	}
}

func GetOpenVPNGroupService(c *gin.Context, id int) {
	var result response.OpenVPNGroup
	iKuai, url, err := iKuaiRequest("/ip-objects/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	_, err = iKuai.
		Get(url)
	if err != nil {
		sendError(c, err)
	} else {
		sendI18n(c, 200, "openvpn.group.get_success", result.Results)
	}
}

func CreateOpenVPNGroupService(c *gin.Context, req request.CreateOpenVPNGroup) {
	var result response.OpenVPNGroupChange
	iKuai, url, err := iKuaiRequest("/ip-objects", &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err := iKuai.
			SetBody(H{
				"group_name":  req.Name,
				"group_value": req.Value,
			}).
			Post(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.group.create_success", result)
		}
	}
}

func UpdateOpenVPNGroupService(c *gin.Context, id int, req request.UpdateOpenVPNGroup) {
	var result response.OpenVPNGroupChange
	iKuai, url, err := iKuaiRequest("/ip-objects/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err := iKuai.
			SetBody(H{
				"group_name":  req.Name,
				"group_value": req.Value,
			}).
			Put(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.group.update_success", result)
		}
	}
}

func DeleteOpenVPNGroupService(c *gin.Context, id int) {
	var result response.OpenVPNGroupChange
	iKuai, url, err := iKuaiRequest("/ip-objects/"+strconv.Itoa(id), &result)
	if err != nil {
		sendError(c, err)
	}
	if nw(c) {
		_, err = iKuai.Delete(url)
		if err != nil {
			sendError(c, err)
		} else {
			sendI18n(c, 200, "openvpn.group.delete_success", result)
		}
	}
}
