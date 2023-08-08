package util

import (
	"math/rand"
	"time"
)

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
