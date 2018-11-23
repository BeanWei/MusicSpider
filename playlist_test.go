package musicspider

import (
	"fmt"
	"testing"
)

const testAll4playlist, index4playlist = false, 2

func Test_playlist(t *testing.T) {

	playlistID := map[string]string{
		"netease": "436843836",
		"tencent": "1144188779",
		"xiami":   "277809522",
		"kugou":   "119859",
		"baidu":   "364201689"}

	if testAll4playlist {
		for _, s := range sites {
			resp := Playlist(s, playlistID[s])
			fmt.Println(resp)
		}
	} else {
		resp := Playlist(sites[index4playlist], playlistID[sites[index4playlist]])
		fmt.Println(resp)
	}

}

/*
TODO: 虾米音乐歌单接口待分析
*/
