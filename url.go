package musicspider

import (
	"fmt"
)

func url(site, id string) map[string]interface{} {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/song/enhance/player/url"
		br := 320 * 1000
		data := fmt.Sprintf(`{"ids": %s, "br": %d}`, id, br)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
		platform, format := "yqq", "json"
		data := fmt.Sprintf(`{"songmid": %s, "platform": %s, "format": %s}`, id, platform, format)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.songservice.getsongs/1.0/"
		data := fmt.Sprintf(`{"songIds": %s}`, id)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "POST"
		url := "http://media.store.kugou.com/v1/get_res_privilege"
		relate, userid, vip, appid, token, behavior, area_code, clientver, rid, rtype := 1, 0, 0, 1000, "", "download", 1, 8990, 0, "audio"
		data := fmt.Sprintf(`{"hash": %s, "relate": %d, "userid": %d, "vip": %d, "appid": %d, "token": %s, "behavior": %s, "area_code": %d, "clientver": %d, "id": %d, "type": %s}`,
										id, relate, userid, vip, appid, token, behavior, area_code, clientver, rid, rtype)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, res, platform, version := "qianqianmini", "baidu.ting.song.getInfos", 1, "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": %s, "from": %s, "method": "%s", "res": %d, "platform": %s, "version": %s}`,
			id, from, method, res, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]interface{}{"status": "404", "result": "暂不支持此站点"}
	}
}
