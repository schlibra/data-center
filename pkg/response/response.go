package response

import (
	"net/http"

	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

// Response 通用响应体（用于 Swagger 文档）
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendJson(c *gin.Context, code int, message string, data any) {
	if !c.Writer.Written() {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		})
	}
}
func SendI18n(c *gin.Context, code int, messageId string, data any) {
	SendJson(c, code, i18n.MustGetMessage(c, messageId), data)
}
func SendError(c *gin.Context, err error) {
	SendJson(c, 500, err.Error(), nil)
}
