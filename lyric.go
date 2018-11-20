package musicspider

import (
	"encoding/base64"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func lyric(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/song/lyric"
		os, lv, kv, tv := "linux", "-1", "-1", "-1"
		data := fmt.Sprintf(`{"id": %s, "os": %s, "lv": %s, "kv": %s, "tv": %s}`,
			id, os, lv, kv, tv)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg"
		g_tk := "5381"
		data := fmt.Sprintf(`{"songmid": %s, "g_tk": %s}`, id, g_tk)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.lyricservice.getsonglyrics/1.0/"
		data := fmt.Sprintf(`{"songId": %s}`, id)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "GET"
		url := "http://krcs.kugou.com/search"
		keyword, ver, client, man := "%20-%20", 1, "mobi", "yes"
		data := fmt.Sprintf(`{"hash": %s, "keyword": %s, "ver": %d, "client": %s, "man": %s}`,
			id, keyword, ver, client, man)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, platform, version := "qianqianmini", "baidu.ting.song.lry", "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": %s, "from": %s, "method": "%s", "platform": %s, "version": %s}`,
			id, from, method, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}

func neteaseLyric(result string) string {
	lyric := gojsonq.New().JSONString(result).Find("lrc.lyric")
	if lyric == nil {
		lyric = ""
	}
	tlyric := gojsonq.New().JSONString(result).Find("tlyric.lyric")
	if tlyric == nil {
		tlyric = ""
	}
	lyric, tlyric = fmt.Sprintf("%#", lyric), fmt.Sprintf("%#", tlyric)
	return fmt.Sprintf(`{"lyric": %s, "tlyric": %s}`, lyric, tlyric)
}

func tencentLyric(result string) string {
	// TODO: 获取完整的json返回
	var lyricStr, tlyricStr string
	lyric := gojsonq.New().JSONString(result).Find("tlyric")
	if lyric == nil {
		lyricStr = ""
	} else {
		lyricEncode := fmt.Sprintf("%#", lyric)
		lyricByte, err := base64.StdEncoding.DecodeString(lyricEncode)
		if err == nil {
			lyricStr = string(lyricByte)
		} else {
			lyricStr = ""
		}
	}
	tlyric := gojsonq.New().JSONString(result).Find("trans")
	if tlyric == nil {
		tlyricStr = ""
	} else {
		tlyricEncode := fmt.Sprintf("%#", lyric)
		tlyricByte, err := base64.StdEncoding.DecodeString(tlyricEncode)
		if err == nil {
			tlyricStr = string(tlyricByte)
		} else {
			tlyricStr = ""
		}
	}
	return fmt.Sprintf(`{"lyric": %s, "tlyric": %s}`, lyricStr, tlyricStr)
}

func xiamiLyric(result string) string {
	return fmt.Sprintf(`{"lyric": %s, "tlyric": %s}`, "", "")
}

func kugouLyric(result string) string {
	accesskey := fmt.Sprintf("%#", gojsonq.New().JSONString(result).Find("candidates.[0].accesskey"))
	id := fmt.Sprintf("%#", gojsonq.New().JSONString(result).Find("candidates.[0].id"))
	reqMethod := "GET"
	url := "http://lyrics.kugou.com/download"
	charset, client, _fmt, ver := "utf-8", "mobi", "lrc", 1
	data := fmt.Sprintf(`{"charset: %s", "accesskey": %s, "id": %s, "client": %s, "fmt", %s, "ver": %d}`, charset, accesskey, id, client, _fmt, ver)
	resp := reqHandler("kugou", reqMethod, url, data)
	res := resp["result"]
	var lyricStr string
	lyric := gojsonq.New().JSONString(res).Find("content")
	if lyric == nil {
		lyricStr = ""
	} else {
		lyricEncode := fmt.Sprintf("%#", lyric)
		lyricByte, err := base64.StdEncoding.DecodeString(lyricEncode)
		if err == nil {
			lyricStr = string(lyricByte)
		} else {
			lyricStr = ""
		}
	}
	return fmt.Sprintf(`{"lyric": %s, "tlyric": %s}`, lyricStr, "")
}

func baiduLyric(result string) string {
	var lyricStr string
	lyric := gojsonq.New().JSONString(result).Find("lrcContent")
	if lyric == nil {
		lyricStr = ""
	} else {
		lyricStr = fmt.Sprintf("%#", lyric)
	}
	return fmt.Sprintf(`{"lyric": %s, "tlyric": %s}`, lyricStr, "")
}
