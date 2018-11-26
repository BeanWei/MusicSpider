package musicspider

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

var sites = [5]string{"netease", "tencent", "xiami", "kugou", "baidu"}

func checkDicKey(dict map[string]int, thisKey ...string) bool {
	if dict == nil {
		return false
	}
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

// MD5加密32位
func md5Encrpyt(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}

// MD5加密16位
func md5Encrpyt16(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	md5hex := hex.EncodeToString(hasher.Sum(nil))
	loop := len(md5hex) / 2
	buf := bytes.NewBufferString("")
	for i := 0; i < loop; i++ {
		idx, _ := strconv.ParseInt(md5hex[i*2:i*2+2], 16, 10)
		buf.WriteByte(byte(int8(idx)))
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// URL Decode
func urlDecode(encodeURL string) string {
	result, _ := url.ParseQuery(encodeURL)
	keys := []string{}
	for key, _ := range result {
		keys = append(keys, key)
	}
	res := strings.Replace(keys[0], "^", "0", -1)
	return res
}
