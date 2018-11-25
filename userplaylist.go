package musicspider

import "fmt"

func UserPlaylist(site, id string, option map[string]int) map[string]string {
	var (
		limit  = 30
		offset = 0
	)
	switch site {
	case "netease":
		reqMethod := "GET"
		url := "http://music.163.com/api/user/playlist/"
		if checkDicKey(option, "offset") {
			offset = option["offset"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		data := fmt.Sprintf(`{"offset": "%d", "limit": "%d", "uid": %s}`, offset, limit, id)
		return reqHandler("netease", reqMethod, url, data)
	default:
		//TODO: 支持其他站点
		return nil

	}
}
