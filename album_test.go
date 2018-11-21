package musicspider

import (
	"fmt"
	"testing"
)

const testAll4album, index4album = false, 0

func Test_album(t *testing.T) {

	albumID := map[string]string{
		"netease": "",
		"tencent": "",
		"xiami":   "",
		"kugou":   "",
		"baidu":   ""}

	if testAll4album {
		for _, s := range sites {
			resp := song(s, albumID[s])
			fmt.Println(resp)
		}
	} else {
		resp := song(sites[index4album], albumID[sites[index4album]])
		fmt.Println(resp)
	}

}
