package musicspider

import (
	"fmt"
	"testing"
)

const testAll4lyric, index4lyric = false, 0

func Test_lyric(t *testing.T) {

	lyricID := map[string]string{
		"netease": "418603133",
		"tencent": "000bdxcy3ta8Q3",
		"xiami":   "1807124714",
		"kugou":   "CB7EE97F4CC11C4EA7A1FA4B516A5D97",
		"baidu":   "578055564"}

	if testAll4lyric {
		for _, s := range sites {
			resp := lyric(s, lyricID[s])
			fmt.Println(resp)
		}
	} else {
		resp := lyric(sites[index4lyric], lyricID[sites[index4lyric]])
		fmt.Println(resp)
	}

}
