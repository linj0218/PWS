package util

import (
	"math/rand"
	"time"
)

// RandStr 产生随机字符串
func RandStr(leng int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandNum 产生随机数
func RandNum(i int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(i)
}
