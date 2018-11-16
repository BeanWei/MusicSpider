package musicspider

import (
	"fmt"
)

func lyric(site, id string) map[string]interface{} {
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
		return map[string]interface{}{"status": "404", "result": "暂不支持此站点"}
	}
}
