package main

import (
	"github.com/gin-gonic/gin"
	"kok/common"
)

func main() {
	db := common.InitDB() //连接数据库
	sqldb, _ := db.DB()   //实例化DB()函数
	defer sqldb.Close()   //延迟关闭数据库连接
	r := gin.Default()    //默认路由分组
	r = CollectRoute(r)   //路由对象更新
	panic(r.Run())        // 监听并在 0.0.0.0:8080 上启动服务
}
