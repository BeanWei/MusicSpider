package musicspider

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"math/rand"
)

func reqHandler(site, reqmethod, url, data string) map[string]interface{} {
	client := &http.Client{}
	dataByte := bytes.NewBuffer([]byte(data))
	req, _ := http.NewRequest(reqmethod, url, dataByte)
	switch site {
	case "netease":
		req.Header.Set("Referer"         , "https://music.163.com/")
		req.Header.Set("Cookie"          , "appver=1.5.9; os=osx; __remember_me=true; osver=%E7%89%88%E6%9C%AC%2010.13.5%EF%BC%88%E7%89%88%E5%8F%B7%2017F77%EF%BC%89")
		req.Header.Set("User-Agent"      , "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko)")
		req.Header.Set("X-Real-IP"       , long2ip(randRange(1884815360, 1884890111)))
		req.Header.Set("Accept"          , "*/*")
		req.Header.Set("Accept-Language" , "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
		req.Header.Set("Connection"      , "keep-alive")
		req.Header.Set("Content-Type"    , "application/x-www-form-urlencoded")
	case "tencent":
		req.Header.Set("Referer"         , "http://y.qq.com")
		req.Header.Set("Cookie"          , "pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1")
		req.Header.Set("User-Agent"      , "QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)")
		req.Header.Set("Accept"          , "*/*")
		req.Header.Set("Accept-Language" , "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
		req.Header.Set("Connection"      , "keep-alive")
		req.Header.Set("Content-Type"    , "application/x-www-form-urlencoded")
	case "xiami":
		req.Header.Set("Cookie"          , "_m_h5_tk=15d3402511a022796d88b249f83fb968_1511163656929; _m_h5_tk_enc=b6b3e64d81dae577fc314b5c5692df3c")
		req.Header.Set("User-Agent"      , "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) XIAMI-MUSIC/3.1.1 Chrome/56.0.2924.87 Electron/1.6.11 Safari/537.36")
		req.Header.Set("Accept"          , "application/json")
		req.Header.Set("Accept-Language" , "zh-CN")
		req.Header.Set("Content-Type"    , "application/x-www-form-urlencoded")
	case "kugou":
		req.Header.Set("User-Agent"      ,"IPhone-8990-searchSong")
		req.Header.Set("UNI-UserAgent"	  ,"iOS11.4-Phone8990-1009-0-WiFi")
	case "baidu":
		req.Header.Set("Cookie"          , "BAIDUID='.$this->getRandomHex(32).':FG=1")
		req.Header.Set("User-Agent"      , "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) baidu-music/1.2.1 Chrome/66.0.3359.181 Electron/3.0.5 Safari/537.36")
		req.Header.Set("Accept"          , "*/*")
		req.Header.Set("Accept-Language" , "zh-CN")
		req.Header.Set("Content-Type"    , "application/json;charset=UTF-8")
	default:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, _ := client.Do(req)
	status := strconv.Itoa(resp.StatusCode)
	if status == "200" {
		body, _ := ioutil.ReadAll(resp.Body)
		feedback := map[string]interface{}{"status": status, "result": body}
		return feedback
	}
	feedback := map[string]interface{}{"status": status, "result": ""}
	return feedback
}

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
		(ip >> 24) & 0xFF,
		(ip >> 16) & 0xFF,
		(ip >> 8) & 0xFF,
		ip & 0xFF)
}

// 字符串反转
func reverseStr(s string) string  {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to -1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}





































