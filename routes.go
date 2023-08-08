package main

import (
	"github.com/gin-gonic/gin"
	"kok/controller"
	"kok/middleware"
)

// gin框架路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", controller.Register)                     //用户注册
	r.POST("/auth/login", controller.Login)                           //用户登录
	r.GET("/auth/info", middleware.AuthMiddleware(), controller.Info) //获取用户信息
	return r
}
