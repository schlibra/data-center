package handler

import (
	"data-center/internal/models/request"
	"data-center/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.create")
	var req request.FrpTokenCreate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	service.CreateFrpTokenService(c, req, row)
}

func ListFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.get")
	service.ListFrpTokenService(c, row)
}

func GetFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.get")
	var req request.FrpTokenInfo
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GetFrpTokenService(c, req, row)
}

func UpdateFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.update")
	var req request.FrpTokenUpdate
	if err := c.ShouldBind(&req); err != nil {
		sendError(c, err)
	}
	var reqId request.FrpTokenUpdate
	if err := c.ShouldBindUri(&reqId); err != nil {
		sendError(c, err)
	}
	req.ID = reqId.ID
	service.UpdateFrpTokenService(c, req, row)
}

func DeleteFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.delete")
	var req request.FrpTokenDelete
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.DeleteFrpTokenService(c, req, row)
}

func GenerateFrpTokenHandler(c *gin.Context) {
	row := parseToken(c)
	checkPermission(c, "frp.token.generate")
	var req request.FrpTokenGenerate
	if err := c.ShouldBindUri(&req); err != nil {
		sendError(c, err)
	}
	service.GenerateFrpTokenService(c, req, row)
}
