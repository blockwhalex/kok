package common

import (
	"github.com/dgrijalva/jwt-go"
	"kok/model"
	"time"
)

// 创建一个用于生成密钥的参数
var jwtkey = []byte("a_secret_crect")

// 定义结构体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成token函数，传入用户名后返回token和err
func ReleaseToken(user model.User) (string, error) {
	expriationTime := time.Now().Add(7 * 24 * time.Hour) //token将在一周后失效
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expriationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //签发时间
			Issuer:    "kok.nomarl",          //签发机构
			Subject:   "user token",          //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //使用HS265算法加密
	tokenString, err := token.SignedString(jwtkey)             //使用jwtkey进行签名

	if err != nil {
		return "", err //返回空值和错误信息
	}
	return tokenString, nil //返回token和空值
}

// 解析tokenString返回claims
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, claims, err
}
