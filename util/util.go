package util

import (
	"math/rand"
	"time"
)

// 生成随机字符串的函数
func RandomString(n int) string {
	var letters = []byte("qwertyuiopASDFGHJKLZXCVBNMasdfghjklzxcvbnm") //随机数种子
	result := make([]byte, n)                                          //根据n控制切片长度
	rand.New(rand.NewSource(time.Now().Unix()))                        //时间随机数
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
