package musicspider

import (
	"fmt"
)

func song(site, id string) map[string]interface{} {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/v3/song/detail/"
		v := 0
		data := fmt.Sprintf(`{"id": %s, "v": %d}`, id, v)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
		platform, format := "yqq", "json"
		data := fmt.Sprintf(`{"songmid": %s, "platform": %s, "format": %s}`, id, platform, format)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.songservice.getsongdetail/1.0/"
		data := fmt.Sprintf(`{"songId": %s}`, id)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugo":
		reqMethod := "POST"
		url := "http://m.kugou.com/app/i/getSongInfo.php"
		cmd, from := "playInfo", "mkugou"
		data := fmt.Sprintf(`{"hash": %s, "cmd": %s, "from": %s }`, id, cmd, from)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, res, platform, version := "qianqianmini", "baidu.ting.song.getInfos", 1, "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": %s, "from": %s, "method": %s, "res", %d, "platform": %s, "version": %s}`,
											id, from, method, res, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]interface{}{"status": "404", "result": "暂不支持此站点"}
	}
}