package musicspider

import (
	"fmt"
	"testing"
)

const testAll4playlist, index4playlist = false, 0

func Test_playlist(t *testing.T) {

	playlistID := map[string]string{
		"netease": "64016",
		"tencent": "1144188779",
		"xiami":   "277809522",
		"kugou":   "119859",
		"baidu":   "364201689"}

	if testAll4playlist {
		for _, s := range sites {
			if s == "xiami" {
				fmt.Println("TODO: 虾米接口待完善")
			} else {
				resp := Playlist(s, playlistID[s])
				fmt.Println(resp)
			}
		}
	} else {
		resp := Playlist(sites[index4playlist], playlistID[sites[index4playlist]])
		fmt.Println(resp)
	}

}

/*
TODO: 虾米音乐歌单接口待分析
*/
