package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

func main() {
	db := InitDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()
	r := gin.Default()
	r.POST("/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		log.Print(name + "," + telephone + "," + password)
		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) <= 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
			return
		}
		//如果用户名为空，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		//判断手机号码是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在！"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		//返回结果

		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

// 生成随机字符串的函数
func RandomString(n int) string {
	var letters = []byte("qwertyuiopASDFGHJKLZXCVBNMasdfghjklzxcvbnm")
	result := make([]byte, n)
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 数据库连接池
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "kok"
	username := "root"
	password := "lomi1015."
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.New(mysql.Config{DriverName: driverName, DSN: args}), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

// 连接数据库判断手机号码是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("Telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
