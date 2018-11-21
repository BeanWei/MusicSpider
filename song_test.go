package musicspider

import (
	"fmt"
	"testing"
)

const testAll4song, index4song = false, 4

func Test_song(t *testing.T) {

	songID := map[string]string{
		"netease": "35847388",
		"tencent": "000bdxcy3ta8Q3",
		"xiami":   "1807124714",
		"kugou":   "d353b69a3b7f1a250000c5daabb8a4f1",
		"baidu":   "578055564"}

	if testAll4song {
		for _, s := range sites {
			resp := song(s, songID[s])
			fmt.Println(resp)
		}
	} else {
		resp := song(sites[index4song], songID[sites[index4song]])
		fmt.Println(resp)
	}

}
