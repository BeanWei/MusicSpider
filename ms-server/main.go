package main

import (
	"github.com/gin-gonic/gin"
	"strings"

	ms "github.com/BeanWei/MusicSpider"
)

/*
localhost:8080/api/v1/action?sites=&keyword
action: search, album, artist, playlist, song, lyric, download
sites: netease, tencent, xiami, kugou, baidu
keyword: ID or keywords
*/

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// api group v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/search/:site/:key", search)
		v1.GET("/album/:site/:key", album)
		v1.GET("/artist/:site/:key", artist)
		v1.GET("/playlist/:site/:key", playlist)
		v1.GET("/song/:site/:key", song)
		v1.GET("/lyric/:site/:key", lyric)
		v1.GET("/download/:site/:key", downloadurl)

	}

	router.Run(":8080")
}

/*=======================Controllers=======================*/

var sites = []string{"netease", "tencent", "xiami", "kugou", "baidu"}

func check(site, key string) (bool, string) {
	site, key = strings.TrimSpace(site), strings.TrimSpace(key)
	if site == "" || key == "" {
		return false, "参数不完整"
	}
	for _, s := range sites {
		if s == site {
			return true, "ok"
		}
	}
	return false, "暂不支持此音乐站点"
}

func search(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Search(site, key, nil)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.SearchFormat(site, feedBack["result"])
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}

func album(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Album(site, key)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.AlbumFormat(site, feedBack["result"])
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}

func artist(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Artist(site, key)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.ArtistFormat(site, feedBack["result"])
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}

func playlist(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Playlist(site, key)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.PlaylistFormat(site, feedBack["result"])
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}

func song(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Song(site, key)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.SongFormat(site, feedBack["result"])
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}

func lyric(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	mapRes := ms.Lyric(site, key)
	c.JSON(200, gin.H{"status": "ok", "result": mapRes})
}

func downloadurl(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	mapRes := ms.Downloadurl(site, key)
	c.JSON(200, gin.H{"status": "ok", "result": mapRes})
}
