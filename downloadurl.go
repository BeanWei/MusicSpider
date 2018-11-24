package musicspider

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"regexp"
)

func Downloadurl(site, id string) map[string]string {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/song/enhance/player/url"
		br := "320000"
		data := fmt.Sprintf(`{"ids": "%s", "br": "%s"}`, id, br)
		return reqHandler("netease", reqMethod, url, data, true)
	case "tencent":
		reqMethod := "GET"
		url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
		platform, format := "yqq", "json"
		data := fmt.Sprintf(`{"songmid": "%s", "platform": "%s", "format": "%s"}`, id, platform, format)
		reqResp := reqHandler("tencent", reqMethod, url, data)
		return tencentURL(reqResp["result"])
	case "xiami":
		// 虾米音乐的直链获取比较特殊, 根据歌曲ID请求一个页面得到经过 凯撒加密 过的下载链接, 分析此凯撒矩阵即可直接获取直链
		reqMethod := "GET"
		//url := "http://h5api.m.xiami.com/h5/mtop.alimusic.music.songservice.getsongs/1.0/"
		url := fmt.Sprintf("http://www.xiami.com/widget/xml-single/uid/0/sid/%s", id)
		htmlText := reqHandler("xiami", reqMethod, url, "")["result"]
		if htmlText == "" {
			return map[string]string{"url": ""}
		}
		reg, _ := regexp.Compile(`<location><!\[CDATA\[(.*?)\]\]></location>`)
		encodeText := fmt.Sprintf("%q", reg.FindStringSubmatch(htmlText)[1])
		downloadUrl := urlDecode(Caesar(encodeText))
		return map[string]string{"url": downloadUrl}
	case "kugou":
		reqMethod := "POST"
		url := "http://media.store.kugou.com/v1/get_res_privilege"
		relate, userid, vip, appid, token, behavior, area_code, clientver, rid, rtype := 1, 0, 0, 1000, "", "download", 1, 8990, 0, "audio"
		data := fmt.Sprintf(`{"relate":"%d","userid":"%d","vip":"%d","appid":"%d","token":"%s","behavior":"%s","area_code":"%d","clientver":"%d","resource":[{"id":%d,"type":"%s","hash":"%s"}]}`,
			relate, userid, vip, appid, token, behavior, area_code, clientver, rid, rtype, id)
		reqResp := reqHandler("kugou", reqMethod, url, data)
		return kugouURL(reqResp["result"])
	case "baidu":
		reqMethod := "GET"
		url := "http://musicapi.taihe.com/v1/restserver/ting"
		e := BaiduAESCBC(id)
		from, method, res, platform, version := "qianqianmini", "baidu.ting.song.getInfos", 1, "darwin", "1.0.0"
		data := fmt.Sprintf(`{"songid": "%s", "from": "%s", "method": "%s", "res": "%d", "platform": "%s", "version": "%s", "e": "%s"}`,
			id, from, method, res, platform, version, e)
		reqResp := reqHandler("baidu", reqMethod, url, data)
		return baiduURL(reqResp["result"])
	default:
		return map[string]string{"status": "404", "result": "暂不支持此站点"}
	}
}

func neteaseURL(result string) map[string]string {
	//jsonDom := gojsonq.New().JSONString(result)
	url := gojsonq.New().JSONString(result).Find("data.[0].uf.url")
	if url == nil {
		if url = gojsonq.New().JSONString(result).Find("data.[0].url"); url == nil {
			return map[string]string{"url": ""}
		}
	}
	return map[string]string{"url": url.(string)}
}

func tencentURL(result string) map[string]string {
	const guid = "2088072923"
	const key = "159C1FE19818BFE9B617AF6A1CFB564526A5290B0E3F3D389DC32A498CAB0B69B16CB7D6D10423AF752CCEF66FD330A36AAC7E812CB30BD2"
	media_mid := gojsonq.New().JSONString(result).Find("data.[0].file.media_mid")
	if media_mid == nil {
		return map[string]string{"url": ""}
	}
	url := fmt.Sprintf("https://dl.stream.qqmusic.qq.com/M800%s.mp3?vkey=%s&guid=%s&uid=0&fromtag=30", media_mid, key, guid)
	return map[string]string{"url": url}
}

func kugouURL(result string) map[string]string {
	hashStr := gojsonq.New().JSONString(result).Find("data.[0].relate_goods.[0].hash").(string)
	reqMethod := "GET"
	req4fownurl := "http://trackercdn.kugou.com/i/v2/"
	key, pid, behavior, cmd, version := md5Encrpyt(hashStr+"kgcloudv2"), 3, "play", 25, 8990
	data := fmt.Sprintf(`{"hash": "%s", "key": "%s", "pid": "%d", "behavior": "%s", "cmd": "%d", "version": "%d"}`, hashStr, key, pid, behavior, cmd, version)
	strResp := reqHandler("kugou", reqMethod, req4fownurl, data)
	res := strResp["result"]
	if res == "" {
		return map[string]string{"url": ""}
	}
	url := gojsonq.New().JSONString(res).Find("url[0]")
	if url == nil {
		return map[string]string{"url": ""}
	}
	return map[string]string{"url": url.(string)}
}

func baiduURL(result string) map[string]string {
	var downloadLink string
	for i := 0; i >= 0; i++ {
		path := fmt.Sprintf("songurl.url.[%d].file_link", i)
		matchText := gojsonq.New().JSONString(result).Find(path)
		if i == 0 {
			if matchText == nil || matchText == "" {
				downloadLink = ""
				break
			}
		} else {
			if matchText == nil || matchText == "" {
				break
			}
		}
		downloadLink = matchText.(string)
	}
	return map[string]string{"url": downloadLink}
}
