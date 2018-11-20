package musicspider

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func checkDicKey(dict map[string]int, thisKey ...string) bool {
	for _, k := range thisKey {
		_, ok := dict[k]
		if !ok {
			return ok
		}
	}
	return true
}

func randRange(min, max int64) int64 {
	diff := max - min
	move := rand.Int63n(diff)
	randNum := min + move
	return randNum
}

// Go版 IP2LONG, LONG2IP => https://blog.csdn.net/zengming00/article/details/80354248
func long2ip(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		(ip>>24)&0xFF,
		(ip>>16)&0xFF,
		(ip>>8)&0xFF,
		ip&0xFF)
}

// 字符串反转
func reverseStr(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// MD5加密
func md5Encrpyt(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
