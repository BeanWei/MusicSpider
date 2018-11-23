package musicspider

import (
	"fmt"
)

func Artist(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := fmt.Sprintf("http://music.163.com/api/v1/artist/%s", id)
		ext, private_cloud, top := "true", "true", 50
		data := fmt.Sprintf(`{"albumId": "%s", "ext": "%s", "private_cloud": "%s", "top": "%d""}`,
			id, ext, private_cloud, top)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_v8_singer_track_cp.fcg"
		begin, num, order, platform, newsong := 0, 50, "listen", "mac", 1
		data := fmt.Sprintf(`{"singermid": "%s", "begin": "%d", "num": "%d", "order": "%s", "platform": "%s", "newsong": "%d"}`,
			id, begin, num, order, platform, newsong)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "'http://h5api.m.xiami.com/h5/mtop.alimusic.music.songservice.getartistsongs/1.0/"
		page, pagesize := 1, 50
		data := fmt.Sprintf(`{"artistId": "%s", "page": "%d", "pageSize": "%d"}`, id, page, pagesize)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "GET"
		url := "http://mobilecdn.kugou.com/api/v3/singer/song"
		area_code, page, plat, pagesize, version := 1, 1, 0, 50, 8990
		data := fmt.Sprintf(`{"singerid": "%s", "area_code": "%d", "page": "%d", "plat": "%d", "pagesize": "%d", "version": "%d"}`,
			id, area_code, page, plat, pagesize, version)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, limits, platform, offset, tinguid, version := "qianqianmini", "baidu.ting.artist.getSongList", 50, "darwin", 0, 0, "11.2.1"
		data := fmt.Sprintf(`{"artistid": "%s", "from": "%s", "method": "%s", "limits": "%d", "platform": "%s", "offset": "%d", "tinguid": "%d", "version": "%s"}`,
			id, from, method, limits, platform, offset, tinguid, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}
