package musicspider

import (
	"bytes"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 请求处理函数，所有的请求均通过此入口
// TODO: 更优雅的处理可变参数 args
// args 为一个bool值列表,
// 列表0为为是否需要加密参数处理的Bool值, 默认不处理加密
func reqHandler(site, reqmethod, url, data string, args ...bool) map[string]string {
	client := http.DefaultClient
	var dataByte io.Reader
	if data == "" {
		dataByte = nil
	} else {
		if len(args) > 0 && args[0] == true {
			if site == "netease" {
				dataByte = NeteaseAESCBC(data)
			}
		} else {
			dataByte = bytes.NewBuffer([]byte(data))
		}
	}
	fmt.Println("the URL is: ", url)
	fmt.Println("the Params is: ", dataByte)
	req, err := RequestHandler(reqmethod, url, dataByte)
	if err != nil {
		fmt.Println(err)
	}
	switch site {
	case "netease":
		req.Header.Set("Host", "music.163.com")
		req.Header.Set("Referer", "http://music.163.com")
		//req.Header.Set("Cookie", "appver=1.5.9; os=osx; __remember_me=true; osver=%E7%89%88%E6%9C%AC%2010.13.5%EF%BC%88%E7%89%88%E5%8F%B7%2017F77%EF%BC%89")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko)")
		req.Header.Set("X-Real-IP", long2ip(randRange(1884815360, 1884890111)))
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "tencent":
		req.Header.Set("Referer", "http://y.qq.com")
		//req.Header.Set("Cookie", "pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1")
		req.Header.Set("User-Agent", "QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "xiami":
		req.Header.Set("Referer", "http://h.xiami.com/")
		req.Header.Set("Cookie", "_m_h5_tk=15d3402511a022796d88b249f83fb968_1511163656929; _m_h5_tk_enc=b6b3e64d81dae577fc314b5c5692df3c")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) XIAMI-MUSIC/3.1.1 Chrome/56.0.2924.87 Electron/1.6.11 Safari/537.36")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "zh-CN")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "kugou":
		req.Header.Set("User-Agent", "IPhone-8990-searchSong")
		req.Header.Set("UNI-UserAgent", "iOS11.4-Phone8990-1009-0-WiFi")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "baidu":
		req.Header.Set("Cookie", "BAIDUID='.$this->getRandomHex(32).':FG=1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) baidu-music/1.2.1 Chrome/66.0.3359.181 Electron/3.0.5 Safari/537.36")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "zh-CN")
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	default:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	fmt.Println("the ReqURL is: ", req.URL.String())
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	status := strconv.Itoa(resp.StatusCode)
	if status == "200" {
		body, _ := ioutil.ReadAll(resp.Body)
		if site == "netease" {
			body, _ = GzipDecode(body)
		}
		feedback := map[string]string{"status": status, "result": string(body)}
		return feedback
	}
	feedback := map[string]string{"status": status, "result": ""}
	return feedback
}

// 对 http.NewRequest 进行封装 自动处理 Get 请求下的参数处理
func RequestHandler(method, url string, body io.Reader) (*http.Request, error) {
	// 对GET请求下传过来的的参数进行处理
	if method == "GET" && body != nil {
		gr := gojsonq.New().Reader(body).Get()
		params := "?"
		for key, value := range gr.(map[string]interface{}) {
			params += fmt.Sprintf("%s=%s&", key, value)
		}
		params = strings.TrimRight(params, "&")
		url += params
	}
	req, err := http.NewRequest(method, url, body)
	return req, err
}
