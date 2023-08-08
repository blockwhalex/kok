package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Rseponse(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// 请求成功
func Success(c *gin.Context, data gin.H, msg string) {
	Rseponse(c, http.StatusOK, 200, data, msg)
}

// 请求失败
func Fail(c *gin.Context, data gin.H, msg string) {
	Rseponse(c, http.StatusOK, 400, data, msg)
}
