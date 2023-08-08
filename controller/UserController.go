package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kok/common"
	"kok/dto"
	"kok/model"
	"kok/response"
	"kok/util"
	"log"
	"net/http"
)

// 用户注册
func Register(c *gin.Context) {
	//实例化GetDB方法
	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Rseponse(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) <= 6 {
		response.Rseponse(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//如果用户名为空，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机号码是否存在
	if isTelephoneExist(DB, telephone) {
		response.Rseponse(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Rseponse(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	//数据库新增操作
	DB.Create(&newUser)
	//返回结果
	response.Success(c, nil, "注册成功")
}

// 用户登录
func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	log.Print(telephone + "," + password)
	//数据验证
	if len(telephone) != 11 {
		response.Rseponse(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	//手机号是否存在
	var user model.User
	DB.Where("Telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Rseponse(c, http.StatusUnprocessableEntity, 422, nil, "该用户尚未注册")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Rseponse(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Rseponse(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Panicf("token generate error : %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

// 获取用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

// 连接数据库判断手机号码是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("Telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
