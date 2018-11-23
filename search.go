package musicspider

import (
	"fmt"
)

var (
	_type  = 1
	limit  = 30
	page   = 1
	offset = 0
)

func Search(site, keyword string, option map[string]int) map[string]string {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/weapi/search/get"
		total, n := "true", "1000"
		if checkDicKey(option, "type") {
			_type = option["type"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		if checkDicKey(option, "page", "limit") {
			offset = (option["page"] - 1) * option["limit"]
		}
		data := fmt.Sprintf(`{"s": "%s", "offset": "%d", "limit": "%d", "type": "%d", "total": "%s", "n": "%s"}`,
			keyword, offset, limit, _type, total, n)
		return reqHandler("netease", reqMethod, url, data, true)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/soso/fcgi-bin/client_search_cp"
		format, aggr, lossless, cr, new_json := "json", 1, 1, 1, 1
		if checkDicKey(option, "page") {
			page = option["page"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		data := fmt.Sprintf(`{"w": "%s", "format": "%s", "p": "%d", "n": "%d", "aggr": "%d", "lossless": "%d", "cr": "%d", "new_json": "%d"}`,
			keyword, format, page, limit, aggr, lossless, cr, new_json)
		return reqHandler("tencent", reqMethod, url, data)
	case "xiami":
		reqMethod := "GET"
		//url := "http://h5api.m.xiami.com/h5/mtop.alimusic.search.searchservice.searchsongs/1.0/"
		url := "http://api.xiami.com/web"
		if checkDicKey(option, "page") {
			page = option["page"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		v, r, appKey := "2.0", "search/songs", "1"
		data := fmt.Sprintf(`{"key": "%s", "page": "%d", "limit": "%d", "v": "%s", "r": "%s", "app_key": "%s"}`,
			keyword, page, limit, v, r, appKey)
		return reqHandler("xiami", reqMethod, url, data)
	case "kugou":
		reqMethod := "GET"
		url := "http://mobilecdn.kugou.com/api/v3/search/song"
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		if checkDicKey(option, "page") {
			page = option["page"]
		}
		api_ver, area_code, correct, plat, tag, sver, showtype, version := 1, 1, 1, 2, 1, 5, 10, 8990
		data := fmt.Sprintf(`{"keyword": "%s", "pagesize": "%d", "page": "%d", "api_ver": "%d", "area_code": "%d", "correct": "%d", "plat": "%d", "tag": "%d", "sver": "%d", "showtype": "%d", "version": "%d"}`,
			keyword, limit, page, api_ver, area_code, correct, plat, tag, sver, showtype, version)
		return reqHandler("kugou", reqMethod, url, data)
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		if checkDicKey(option, "page") {
			page = option["page"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		from, method, isNew, platform, version := "qianqianmini", "baidu.ting.search.merge", 1, "darwin", "11.2.1"
		data := fmt.Sprintf(`{"query": "%s", "page_no": "%d", "page_size": "%d", "from": "%s", "method": "%s", "isNew": "%d", "platform": "%s", "version": "%s"}`,
			keyword, page, limit, from, method, isNew, platform, version)
		return reqHandler("baidu", reqMethod, url, data)
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}
