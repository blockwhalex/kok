package main

import (
	"github.com/gin-gonic/gin"
	"kok/common"
)

func main() {
	db := common.InitDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
