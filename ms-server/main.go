package main

import (
	"fmt"
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
		v1.GET("/search/:site/:key", Search)
		/*
			v1.GET("/album/:site/:key")
			v1.GET("/artist/:site/:key")
			v1.GET("/playlist/:site/:key")
			v1.GET("/song/:site/:key")
			v1.GET("/lyric/:site/:key")
			v1.GET("/download/:site/:key")
		*/
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

func Search(c *gin.Context) {
	site, key := c.Param("site"), c.Param("key")
	resBool, resStr := check(site, key)
	if !resBool {
		c.JSON(200, gin.H{"status": "error", "result": resStr})
	}
	feedBack := ms.Search(site, key, nil)
	if feedBack["result"] == "" {
		c.JSON(200, gin.H{"status": "ok", "result": "null"})
	} else {
		result := ms.Format(site, feedBack["result"])
		fmt.Println("Server => the result is: ", result)
		c.JSON(200, gin.H{"status": "ok", "result": result})
	}
}
