package musicspider

import (
	"fmt"
)

func playlist(site, id string) map[string]interface{} {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/v3/playlist/detail"
		s, n, t := 0, 1000, 0
		data := fmt.Sprintf(`{"id": %s, "s": %d, "n": %d, "t": %d}`, id, s, n, t)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_v8_playlist_cp.fcg"
		format, newsong, platform := "json", 1, "jqspaframe.json"
		data := fmt.Sprintf(`{"id": %s, "format": %s, "newsong": %d, "platform": %s}`,
										id, format, newsong, platform)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.list.collectservice.getcollectdetail/1.0/"
		isFullTags, page, pageSize := "false", 1, 1000
		data := fmt.Sprintf(`{"listId": %s, "isFullTags": %s, "page": %d, "pageSize": %d}`,
											id, isFullTags, page, pageSize)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "GET"
		url := "http://mobilecdn.kugou.com/api/v3/special/song"
		area_code, plat, page, pagesize, version := 1, 2, 1, -1, 8990
		data := fmt.Sprintf(`{"specialid": %s, "area_code": %d, "plat": %d, "page": %d, "pagesize": %d, "version": %d}`,
			id, area_code, plat, page, pagesize, version)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, platform, version := "qianqianmini", "baidu.ting.diy.gedanInfo", "darwin", "11.2.1"
		data := fmt.Sprintf(`{"listid": %s, "from": %s, "method": "%s", "platform": %s, "version": %s}`,
			id, from, method, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]interface{}{"status": "404", "result": "暂不支持此站点"}
	}
}
