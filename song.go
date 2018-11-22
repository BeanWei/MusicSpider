package musicspider

import (
	"fmt"
)

func song(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "GET"
		// 这里get请求传递过去的参数不能很好的识别进行拼接，所以直接传一个完整的URl去请求
		url := fmt.Sprintf(`http://music.163.com/api/v3/song/detail?id=%s&c=[{"id":"%s"}]`, id, id)
		return reqHandler("netease", reqMethod, url, "")
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
		platform, format := "yqq", "json"
		data := fmt.Sprintf(`{"songmid": "%s", "platform": "%s", "format": "%s"}`, id, platform, format)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		//url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.songservice.getsongdetail/1.0/"
		url := fmt.Sprintf("http://www.xiami.com/song/playlist/id/%s/object_name/default/object_id/0/cat/json", id)
		return reqHandler("xiami", reqMethod, url, "")
	case "kugou":
		reqMethod := "GET"
		url := "http://m.kugou.com/app/i/getSongInfo.php"
		cmd := "playInfo"
		data := fmt.Sprintf(`{"cmd": "%s", "hash": "%s"}`, cmd, id)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		e := BaiduAESCBC(id)
		from, method, res, platform, version := "qianqianmini", "baidu.ting.song.getInfos", 1, "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": "%s", "from": "%s", "method": "%s", "res": "%d", "platform": "%s", "version": "%s", "e": "%s"}`,
			id, from, method, res, platform, version, e)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}
