package middleware

import (
	"github.com/gin-gonic/gin"
	"kok/common"
	"kok/model"
	"net/http"
	"strings"
)

// 用户认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取 authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "no have power！"})
			c.Abort()
			return
		}
		//取前7位
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token is invalid！"})
			c.Abort()
			return
		}
		//验证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.DB
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "user not exist！"})
			c.Abort()
			return
		}

		//用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
