package musicspider

import (
	"encoding/base64"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"regexp"
)

func lyric(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "GET"
		url := "http://music.163.com/api/song/lyric"
		os, lv, kv, tv := "linux", "-1", "-1", "-1"
		data := fmt.Sprintf(`{"id": "%s", "os": "%s", "lv": "%s", "kv": "%s", "tv": "%s"}`,
			id, os, lv, kv, tv)
		reqResp := reqHandler("netease", reqMethod, url, data)
		return neteaseLyric(reqResp["result"])
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg"
		g_tk := "5381"
		data := fmt.Sprintf(`{"songmid": "%s", "g_tk": %s}`, id, g_tk)
		reqResp := reqHandler("tencent", reqMethod, url, data)
		return tencentLyric(reqResp["result"])
	case "xiami":
		// TODO: 虾米音乐歌词链接暂时只在歌曲信息内找得到，这里只能先请求歌曲信息然后再提取出歌词链接
		/*/
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.lyricservice.getsonglyrics/1.0/"
		data := fmt.Sprintf(`{"songId": %s}`, id)
		return reqHandler("xiami", reqMethod, url, data)
		*/
		songInfo := song("xiami", id)
		lrcURI := gojsonq.New().JSONString(songInfo["result"]).Find("data.trackList.[0].lyric")
		var lyricLink string
		if lrcURI == nil {
			return map[string]string{"lyric": "", "tlyric": ""}
		} else {
			lyricLink = "https:" + lrcURI.(string)
		}
		reqResp := reqHandler("xiami", "GET", lyricLink, "")
		return xiamiLyric(reqResp["result"])
	case "kugou":
		reqMethod := "GET"
		//url := "http://krcs.kugou.com/search"
		url := "http://www.kugou.com/yy/index.php"
		r := "play/getdata"
		data := fmt.Sprintf(`{"r": "%s", "hash": "%s"}`, r, id)
		reqResp := reqHandler("kugou", reqMethod, url, data)
		return kugouLyric(reqResp["result"])
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, platform, version := "qianqianmini", "baidu.ting.song.lry", "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": "%s", "from": "%s", "method": "%s", "platform": "%s", "version": "%s"}`,
			id, from, method, platform, version)
		reqResp := reqHandler("baidu", reqMethod, url, data)
		return baiduLyric(reqResp["result"])
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}

func neteaseLyric(result string) map[string]string {
	var lyricStr, tlyricStr string
	lyric := gojsonq.New().JSONString(result).Find("lrc.lyric")
	if lyric == nil {
		lyricStr = ""
	} else {
		lyricStr = lyric.(string)
	}
	tlyric := gojsonq.New().JSONString(result).Find("tlyric.lyric")
	if tlyric == nil {
		tlyricStr = ""
	} else {
		tlyricStr = tlyric.(string)
	}
	return map[string]string{"lyric": lyricStr, "tlyric": tlyricStr}
}

func tencentLyric(result string) map[string]string {
	var lyricStr, tlyricStr string
	reg, err := regexp.Compile(`MusicJsonCallback\((.*?)\)`)
	if err != nil {
		return map[string]string{"lyric": "", "tlyric": ""}
	}
	jsRes := reg.FindStringSubmatch(result)[1]
	lyric := gojsonq.New().JSONString(jsRes).Find("lyric")
	if lyric == nil || lyric == "" {
		lyricStr = ""
	} else {
		lyricByte, err := base64.StdEncoding.DecodeString(lyric.(string))
		if err == nil {
			lyricStr = string(lyricByte)
		} else {
			lyricStr = ""
		}
	}
	tlyric := gojsonq.New().JSONString(jsRes).Find("trans")
	if tlyric == nil || tlyric == "" {
		tlyricStr = ""
	} else {
		tlyricByte, err := base64.StdEncoding.DecodeString(tlyric.(string))
		if err == nil {
			tlyricStr = string(tlyricByte)
		} else {
			tlyricStr = ""
		}
	}
	return map[string]string{"lyric": lyricStr, "tlyric": tlyricStr}
}

func xiamiLyric(result string) map[string]string {
	// 可能需要扩展虾米歌词格式解析
	return map[string]string{"lyric": result, "tlyric": ""}
}

func kugouLyric(result string) map[string]string {
	var lyricStr string
	lyric := gojsonq.New().JSONString(result).Find("data.lyrics")
	if lyric == nil || lyric == "" {
		lyricStr = ""
	} else {
		lyricStr = lyric.(string)
	}
	return map[string]string{"lyric": lyricStr, "tlyric": ""}
}

func baiduLyric(result string) map[string]string {
	lyric := gojsonq.New().JSONString(result).Find("lrcContent")
	var lyricStr string
	if lyric == nil {
		lyricStr = ""
	} else {
		lyricStr = lyric.(string)
	}
	return map[string]string{"lyric": lyricStr, "tlyric": ""}
}
