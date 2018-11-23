package musicspider

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func safeParse(jsonStr, path string) interface{} {
	result := gojsonq.New().JSONString(jsonStr).Find(path)
	if result == nil || result == "" {
		return ""
	}
	return result
}

type SearchResp struct {
	source string
	count  int
	songs  []SearchSongs
}

type SearchSongs struct {
	songID    string
	songName  string
	singer    SingerMap
	album     AlbumMap
	sourceURL string
}

type SingerMap struct {
	id   string
	name string
}

type AlbumMap struct {
	id   string
	name string
}

func Format(site, jsonStr string) (shr SearchResp) {
	var shs SearchSongs
	var shslist []SearchSongs
	shr.source, shr.count, shr.songs = "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return shr
	}
	switch site {
	case "netease":
		for i := 0; i >= 0; i++ {
			shs.songID = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].id", i)).(string)
			if shs.songID == "" {
				break
			}
			shs.songName = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].name", i)).(string)
			shs.singer.id = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].id", i)).(string)
			shs.singer.name = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].name", i)).(string)
			shs.album.id = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.id", i)).(string)
			shs.album.name = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.name", i)).(string)
			shs.sourceURL = fmt.Sprintf("https://music.163.com/#/song?id=%s", shs.songID)
			shslist = append(shslist, shs)
		}
		shr.source, shr.count, shr.songs = "netease", len(shslist), shslist
	case "tencent":
		for i := 0; i >= 0; i++ {
			shs.songID = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].mid", i)).(string)
			if shs.songID == "" {
				break
			}
			shs.songName = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].title", i)).(string)
			shs.singer.id = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].id", i)).(string)
			shs.singer.name = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].name", i)).(string)
			shs.album.id = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.mid", i)).(string)
			shs.album.name = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.name", i)).(string)
			shs.sourceURL = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", shs.songID)
			shslist = append(shslist, shs)
		}
		shr.source, shr.count, shr.songs = "tencent", len(shslist), shslist
	case "xiami":
		for i := 0; i >= 0; i++ {
			shs.songID = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_id", i)).(string)
			if shs.songID == "" {
				break
			}
			shs.songName = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_name", i)).(string)
			shs.singer.id = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_id", i)).(string)
			shs.singer.name = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_name", i)).(string)
			shs.album.id = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_id", i)).(string)
			shs.album.name = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_name", i)).(string)
			shs.sourceURL = fmt.Sprintf("https://www.xiami.com/song/%s", shs.songID)
			shslist = append(shslist, shs)
		}
		shr.source, shr.count, shr.songs = "xiami", len(shslist), shslist
	case "kugou":
		for i := 0; i >= 0; i++ {
			shs.songID = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i)).(string)
			if shs.songID == "" {
				break
			}
			shs.songName = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].songname", i)).(string)
			shs.singer.id = ""
			shs.singer.name = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].singername", i)).(string)
			shs.album.id = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_id", i)).(string)
			shs.album.name = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_name", i)).(string)
			shs.sourceURL = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", shs.songID)
			shslist = append(shslist, shs)
		}
		shr.source, shr.count, shr.songs = "kugou", len(shslist), shslist
	case "baidu":
		for i := 0; i >= 0; i++ {
			shs.songID = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].song_id", i)).(string)
			if shs.songID == "" {
				break
			}
			shs.songName = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].title", i)).(string)
			shs.singer.id = ""
			shs.singer.name = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].author", i)).(string)
			shs.album.id = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_id", i)).(string)
			shs.album.name = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_title", i)).(string)
			shs.sourceURL = fmt.Sprintf("http://music.taihe.com/song/%s", shs.songID)
			shslist = append(shslist, shs)
		}
		shr.source, shr.count, shr.songs = "baidu", len(shslist), shslist
	default:
	}
	return shr
}
