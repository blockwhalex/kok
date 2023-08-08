package main

import (
	"github.com/gin-gonic/gin"
	"kok/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", controller.Register)
	return r
}
