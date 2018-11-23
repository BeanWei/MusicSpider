package musicspider

import (
	"fmt"
	"testing"
)

const testAll4song, index4song = true, 4

func Test_song(t *testing.T) {

	songID := map[string]string{
		"netease": "35847388",
		"tencent": "000bdxcy3ta8Q3",
		"xiami":   "1807124714",
		"kugou":   "d353b69a3b7f1a250000c5daabb8a4f1",
		"baidu":   "7325804"}

	if testAll4song {
		for _, s := range sites {
			resp := Song(s, songID[s])
			fmt.Println(resp)
		}
	} else {
		resp := Song(sites[index4song], songID[sites[index4song]])
		fmt.Println(resp)
	}

}

/*
TODO: 百度音乐的接口访问会出现IP不被允许或者提示参数错误？ 可能得挂代理防Ban.
*/
