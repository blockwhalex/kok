package main

import (
	"github.com/gin-gonic/gin"
	"kok/controller"
	"kok/middleware"
	"net/http"
)

// gin框架路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", controller.Register)                     //用户注册
	r.POST("/auth/login", controller.Login)                           //用户登录
	r.GET("/auth/info", middleware.AuthMiddleware(), controller.Info) //获取用户信息
	return r
}

// 跨域请求配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 当Access-Control-Allow-Credentials为true时，将*替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 可选
			c.Set("content-type", "application/json")   // 可选
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
