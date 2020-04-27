package utils

import (
	"math/rand"
	"time"
)

// 返回n个随机字符
func RandomString(n int) string {
	var letters = []byte("qwertiyoupanlfmNDOSANFZPJOJFPQWO")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}