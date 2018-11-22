package musicspider

import (
	"fmt"
	"testing"
)

const testAll4album, index4album = false, 0

func Test_album(t *testing.T) {

	albumID := map[string]string{
		"netease": "35501199",
		"tencent": "003rytri2FHG3V",
		"xiami":   "9cGicX47eb8",
		"kugou":   "1645030",
		"baidu":   "541223882"}

	if testAll4album {
		for _, s := range sites {
			resp := album(s, albumID[s])
			fmt.Println(resp)
		}
	} else {
		resp := album(sites[index4album], albumID[sites[index4album]])
		fmt.Println(resp)
	}

}

/*
TODO: 虾米专辑详情接口容易被ban，待完善。
*/
