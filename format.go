package musicspider

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"strings"
)

func safeParse(jsonStr, path string) interface{} {
	result := gojsonq.New().JSONString(jsonStr).Find(path)
	if result == nil || result == "" {
		return ""
	}
	return result
}

/*===================Search Songs API Response Format==================================*/
func SearchFormat(site, jsonStr string) map[string]interface{} {
	shs, shslist, shr := map[string]interface{}{}, []interface{}{}, map[string]interface{}{}
	shr["source"], shr["count"], shr["songs"] = "", 0, nil

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
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].id", i)).(string)
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].name", i)).(string)
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].id", i)).(string)
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].name", i)).(string)
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.id", i)).(string)
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.name", i)).(string)
			shs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
		shr["source"], shr["count"], shr["songs"] = "netease", len(shslist), shslist
	case "tencent":
		for i := 0; i >= 0; i++ {
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].mid", i)).(string)
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].title", i)).(string)
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].id", i)).(string)
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].name", i)).(string)
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.mid", i)).(string)
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.name", i)).(string)
			shs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", shs["song_id"])
			shslist = append(shslist, shs)
		}
		shr["source"], shr["count"], shr["songs"] = "tencent", len(shslist), shslist
	case "xiami":
		for i := 0; i >= 0; i++ {
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_id", i)).(string)
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_name", i)).(string)
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_id", i)).(string)
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_name", i)).(string)
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_id", i)).(string)
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_name", i)).(string)
			shs["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
		shr["source"], shr["count"], shr["songs"] = "xiami", len(shslist), shslist
	case "kugou":
		for i := 0; i >= 0; i++ {
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i)).(string)
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].songname", i)).(string)
			shs["singer_id"] = ""
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].singername", i)).(string)
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_id", i)).(string)
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_name", i)).(string)
			shs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
		shr["source"], shr["count"], shr["songs"] = "kugou", len(shslist), shslist
	case "baidu":
		for i := 0; i >= 0; i++ {
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].song_id", i)).(string)
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].title", i)).(string)
			shs["singer_id"] = ""
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].author", i)).(string)
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_id", i)).(string)
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_title", i)).(string)
			shs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
		shr["source"], shr["count"], shr["songs"] = "baidu", len(shslist), shslist
	default:
	}
	return shr
}

/*===================Album Songs API Response Format==================================*/
func AlbumFormat(site, jsonStr string) map[string]interface{} {
	songs, songslist, album := map[string]interface{}{}, []interface{}{}, map[string]interface{}{}
	album["source"], album["album_name"], album["singer_name"], album["singer_id"], album["brief_desc"], album["publish_time"], album["cover_url"], album["total_songs"], album["songs"] = "", "", "", "", "", "", "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return album
	}
	switch site {
	case "netease":
		album["source"] = "netease"
		album["album_name"] = safeParse(jsonStr, "album.name").(string)
		album["singer_name"] = safeParse(jsonStr, "album.artist.[0].name").(string)
		album["singer_id"] = safeParse(jsonStr, "album.artist.[0].id").(string)
		album["brief_desc"] = safeParse(jsonStr, "album.briefDesc").(string)
		album["publish_time"] = safeParse(jsonStr, "album.publishTime").(string)
		album["cover_url"] = safeParse(jsonStr, "album.blurPicUrl").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "tencent":
		album["source"] = "tencent"
		album["album_name"] = safeParse(jsonStr, "data.getAlbumInfo.Falbum_name").(string)
		album["singer_name"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_name").(string)
		album["singer_id"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_mid").(string)
		album["brief_desc"] = safeParse(jsonStr, "data.getAlbumDesc.Falbum_desc").(string)
		album["publish_time"] = safeParse(jsonStr, "data.getAlbumInfo.Fpublic_time").(string)
		album["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].mid", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "xiami":
		album["source"] = "xiami"
		album["album_name"] = safeParse(jsonStr, "data.[0].album_name").(string)
		album["singer_name"] = safeParse(jsonStr, "data.[0].artist_name").(string)
		album["singer_id"] = safeParse(jsonStr, "data.[0].artist_id").(string)
		album["brief_desc"] = ""
		album["publish_time"] = ""
		// TODO: 获取虾米音乐专辑高清封面
		album["cover_url"] = "https://pic.xiami.net/" + safeParse(jsonStr, "data.[0].album_logo").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.[%d].songId", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].songName", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "kugou":
		album["source"] = "kugou"
		album["album_name"] = ""
		album["singer_name"] = strings.Split(safeParse(jsonStr, "data.info.[0].filename").(string), " - ")[0]
		album["singer_id"] = ""
		album["brief_desc"] = ""
		album["publish_time"] = ""
		album["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i)).(string)
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "baidu":
		album["source"] = "baidu"
		album["album_name"] = safeParse(jsonStr, "albumInfo.title").(string)
		album["singer_name"] = safeParse(jsonStr, "albumInfo.author").(string)
		album["singer_id"] = safeParse(jsonStr, "albumInfo.artist_id").(string)
		album["brief_desc"] = safeParse(jsonStr, "albumInfo.info").(string)
		album["publish_time"] = safeParse(jsonStr, "albumInfo.publishtime").(string)
		album["cover_url"] = safeParse(jsonStr, "albumInfo.pic_s1000").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].song_id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].title", i)).(string)
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	default:
	}
	return album
}

/*===================Artist Songs API Response Format==================================*/
func ArtistFormat(site, jsonStr string) map[string]interface{} {
	songs, songslist, artist := map[string]interface{}{}, []interface{}{}, map[string]interface{}{}
	artist["source"], artist["singer_name"], artist["brief_desc"], artist["cover_url"], artist["total_songs"], artist["songs"] = "", "", "", "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return artist
	}
	switch site {
	case "netease":
		artist["source"] = "netease"
		artist["singer_name"] = safeParse(jsonStr, "artist.name").(string)
		artist["brief_desc"] = safeParse(jsonStr, "artist.briefDesc").(string)
		artist["cover_url"] = safeParse(jsonStr, "artist.picUrl").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("hotSongs.[%d].id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("hotSongs.[%d].name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "tencent":
		artist["source"] = "tencent"
		artist["singer_name"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_name").(string)
		artist["brief_desc"] = safeParse(jsonStr, "data.getAlbumDesc.Falbum_desc").(string)
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.list.[%d].musicData.mid", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].musicData.name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "xiami":
		// TODO: 虾米艺人API接口待完善
		return artist
	case "kugou":
		artist["source"] = "kugou"
		artist["singer_name"] = strings.Split(safeParse(jsonStr, "data.info.[0].filename").(string), " - ")[0]
		artist["brief_desc"] = ""
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i)).(string)
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "baidu":
		artist["source"] = "baidu"
		artist["singer_name"] = safeParse(jsonStr, "songlist.[0].author").(string)
		artist["brief_desc"] = ""
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].song_id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].title", i)).(string)
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	default:
	}
	return artist
}

/*===================Playlist Songs API Response Format==================================*/
func PlaylistFormat(site, jsonStr string) map[string]interface{} {
	songs, songslist, playlist := map[string]interface{}{}, []interface{}{}, map[string]interface{}{}
	playlist["source"], playlist["creator_nickname"], playlist["playlist_name"], playlist["cover_url"], playlist["brief_desc"], playlist["tags"], playlist["creat_time"], playlist["play_count"], playlist["subscribed_count"], playlist["total_songs"], playlist["songs"] = "", "", "", "", "", "", "", "", "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return playlist
	}
	switch site {
	case "netease":
		playlist["source"] = "netease"
		playlist["creator_nickname"] = safeParse(jsonStr, "playlist.creator.nickname").(string)
		playlist["playlist_name"] = safeParse(jsonStr, "playlist.name").(string)
		playlist["brief_desc"] = safeParse(jsonStr, "playlist.description").(string)
		tagsStr := ""
		for i := 0; i >= 0; i++ {
			tag := safeParse(jsonStr, fmt.Sprintf("playlist.tags.[%d]", i)).(string) + ","
			if tag == "" {
				break
			}
			tagsStr += tag
		}
		playlist["tags"] = strings.TrimRight(tagsStr, ",")
		playlist["creat_time"] = safeParse(jsonStr, "playlist.createTime").(string)
		playlist["cover_url"] = safeParse(jsonStr, "playlist.coverImgUrl").(string)
		playlist["play_count"] = safeParse(jsonStr, "playlist.playCount").(string)
		playlist["subscribed_count"] = safeParse(jsonStr, "playlist.subscribedCount").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].name", i)).(string)
			songs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].ar.[0].id", i)).(string)
			songs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].ar.[0].name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["songs"] = len(songslist), songslist
	case "tencent":
		playlist["source"] = "tencent"
		playlist["creator_nickname"] = safeParse(jsonStr, "data.cdlist.nickname").(string)
		playlist["playlist_name"] = safeParse(jsonStr, "data.cdlist.dissname").(string)
		playlist["brief_desc"] = safeParse(jsonStr, "data.cdlist.desc").(string)
		tagsStr := ""
		for i := 0; i >= 0; i++ {
			tag := safeParse(jsonStr, fmt.Sprintf("data.cdlist.tags.[%d].name", i)).(string) + ","
			if tag == "" {
				break
			}
			tagsStr += tag
		}
		playlist["tags"] = strings.TrimRight(tagsStr, ",")
		playlist["creat_time"] = safeParse(jsonStr, "data.cdlist.ctime").(string)
		playlist["cover_url"] = safeParse(jsonStr, "data.cdlist.logo").(string)
		playlist["play_count"] = safeParse(jsonStr, "data.cdlist.visitnum").(string)
		playlist["subscribed_count"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].mid", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].name", i)).(string)
			songs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].singer.mid", i)).(string)
			songs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].singer.name", i)).(string)
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["songs"] = len(songslist), songslist
	case "xiami":
		// TODO: 虾米歌单API接口待完善
		return playlist
	case "kugou":
		playlist["source"] = "kugou"
		playlist["creator_nickname"] = ""
		playlist["playlist_name"] = ""
		playlist["brief_desc"] = ""
		playlist["tags"] = ""
		playlist["creat_time"] = ""
		playlist["cover_url"] = ""
		playlist["play_count"] = ""
		playlist["subscribed_count"] = ""
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i)).(string)
			songs["singer_id"] = ""
			songs["singer_name"] = strings.Split(songs["song_name"].(string), " - ")[0]
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["songs"] = len(songslist), songslist
	case "baidu":
		playlist["source"] = "baidu"
		playlist["creator_nickname"] = ""
		playlist["playlist_name"] = safeParse(jsonStr, "title").(string)
		playlist["brief_desc"] = safeParse(jsonStr, "desc").(string)
		playlist["tags"] = safeParse(jsonStr, "tag").(string)
		playlist["creat_time"] = ""
		playlist["cover_url"] = safeParse(jsonStr, "pic_w700").(string)
		playlist["play_count"] = safeParse(jsonStr, "listenum").(string)
		playlist["subscribed_count"] = safeParse(jsonStr, "collectnum").(string)
		for i := 0; i >= 0; i++ {
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].song_id", i)).(string)
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].title", i)).(string)
			songs["singer_id"] = ""
			songs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].author", i)).(string)
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["songs"] = len(songslist), songslist
	default:
	}
	return playlist
}

/*===================Song Info API Response Format==================================*/
func SongFormat(site, jsonStr string) map[string]interface{} {
	songinfo := map[string]interface{}{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return nil
	}
	switch site {
	case "netease":
		songinfo["source"] = "netease"
		songinfo["song_name"] = safeParse(jsonStr, "songs.[0].name").(string)
		songinfo["singer_id"] = safeParse(jsonStr, "songs.ar.[0].id").(string)
		songinfo["singer_name"] = safeParse(jsonStr, "songs.ar.[0].name").(string)
		songinfo["publish_time"] = safeParse(jsonStr, "songs.publishTime").(string)
		songinfo["cover_url"] = safeParse(jsonStr, "songs.al.picUrl").(string)
		songinfo["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", safeParse(jsonStr, "songs.[0].id").(string))

	case "tencent":
		songinfo["source"] = "tencent"
		songinfo["song_name"] = safeParse(jsonStr, "data.[0].name").(string)
		songinfo["singer_id"] = safeParse(jsonStr, "data.singer.[0].id").(string)
		songinfo["singer_name"] = safeParse(jsonStr, "data.singer.[0].name").(string)
		songinfo["publish_time"] = safeParse(jsonStr, "data.[0].time_public").(string)
		songinfo["cover_url"] = ""
		songinfo["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", safeParse(jsonStr, "data.[0].mid").(string))
	case "xiami":
		songinfo["source"] = "xiami"
		songinfo["song_name"] = safeParse(jsonStr, "data.trackList.[0].songName").(string)
		songinfo["singer_id"] = safeParse(jsonStr, "data.trackList.[0].artistId").(string)
		songinfo["singer_name"] = safeParse(jsonStr, "data.trackList.[0].singers").(string)
		songinfo["publish_time"] = safeParse(jsonStr, "data.trackList.[0].demoCreateTime").(string)
		songinfo["cover_url"] = safeParse(jsonStr, "data.trackList.[0].pic").(string)
		songinfo["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", safeParse(jsonStr, "data.trackList.[0].songId").(string))
	case "kugou":
		songinfo["source"] = "kugou"
		songinfo["song_name"] = safeParse(jsonStr, "songName").(string)
		songinfo["singer_id"] = safeParse(jsonStr, "singerId").(string)
		songinfo["singer_name"] = safeParse(jsonStr, "choricSinger").(string)
		songinfo["publish_time"] = ""
		songinfo["cover_url"] = safeParse(jsonStr, "imgUrl").(string)
		songinfo["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", safeParse(jsonStr, "hash").(string))
	case "baidu":
		songinfo["source"] = "baidu"
		songinfo["song_name"] = safeParse(jsonStr, "songinfo.title").(string)
		songinfo["singer_id"] = safeParse(jsonStr, "songinfo.artist_id").(string)
		songinfo["singer_name"] = safeParse(jsonStr, "songinfo.author").(string)
		songinfo["publish_time"] = safeParse(jsonStr, "songinfo.publishtime").(string)
		songinfo["cover_url"] = safeParse(jsonStr, "songinfo.pic_huge").(string)
		songinfo["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", safeParse(jsonStr, "songinfo.song_id").(string))
	default:
	}
	return songinfo
}
