package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kok/common"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB() //连接数据库
	sqldb, _ := db.DB()   //实例化DB()函数
	defer sqldb.Close()   //延迟关闭数据库连接
	r := gin.Default()    //默认路由分组
	r.Use(Cors())         //开启中间件 允许使用跨域请求
	r = CollectRoute(r)   //路由对象更新
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port)) // 监听并在指定port端口启动服务
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080上启动服务
}

// 读取项目配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read faild to config ,err:" + err.Error())
	}
}
