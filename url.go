package musicspider

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
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

const unexceptResp = `{"url": "", "size": 0, "br": -1}`

func neteaseURL(result string) string {
	//jsonDom := gojsonq.New().JSONString(result)
	url := gojsonq.New().JSONString(result).Find("data.[0].uf.url")
	if url == nil {
		if url = gojsonq.New().JSONString(result).Find("data.[0].url"); url == nil {
			return unexceptResp
		}
	}
	size := gojsonq.New().JSONString(result).Find("data.[0].size")
	br := gojsonq.New().JSONString(result).Find("data.[0].br")
	return fmt.Sprintf(`{"url": %s, "size": %d, "br": %d}`, url, size, br)
}

func tencentURL(result string) string {
	const key  = "58A42A941AEC67EDC9792B5A824F8E8AAB5D3F5641628A002364306E7796FE98CAE45F2F1EEC7BAB5DD6F6C158C08D24471C31CB6200B625"
	if media_mid := gojsonq.New().JSONString(result).Find("data.[0].file.media_mid"); media_mid == nil {
		return unexceptResp
	}
	//https://dl.stream.qqmusic.qq.com/'.$vo[1].$data['data'][0]['file']['media_mid'].'.'.$vo[2].'?vkey='.$key.'&guid='.$guid.'&uid=0&fromtag=30
}















































