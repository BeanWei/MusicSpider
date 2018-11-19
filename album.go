package musicspider

import (
	"fmt"
)

func album(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := fmt.Sprintf("http://music.163.com/api/v1/album/%s", id)
		total, offset, limit, ext, private_cloud := "true", "0", "1000", "true", "true"
		data := fmt.Sprintf(`{"id": %s, "total": %s, "offset": %s, "limit": %s, "ext": %s, "private_cloud": %s }`,
			id, total, offset, limit, ext, private_cloud)
		return reqHandler("netease", reqMethod, url, data)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_v8_album_detail_cp.fcg"
		platform, format, newsong := "mac", "json", 1
		data := fmt.Sprintf(`{"albummid": %s, "platform": %s, "format": %s, "newsong": %d }`,
			id, platform, format, newsong)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.albumservice.getalbumdetail/1.0/"
		data := fmt.Sprintf(`{"albumId": %s}`, id)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "GET"
		url := "http://mobilecdn.kugou.com/api/v3/album/song"
		area_code, plat, page, pagesize, version := 1, 2, 1, -1, 8990
		data := fmt.Sprintf(`{"albumid": %s, "area_code": %d, "plat": %d, "page": %d, "pagesize": %d, "version": %d}`,
			id, area_code, plat, page, pagesize, version)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		from, method, platform, version := "qianqianmini", "baidu.ting.album.getAlbumInfo", "darwin", "11.2.1"
		data := fmt.Sprintf(`{"album_id": %s, "from": %s, "method": "%s", "platform": %s, "version": %s}`,
			id, from, method, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}
