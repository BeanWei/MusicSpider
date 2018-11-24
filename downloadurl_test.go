package musicspider

import (
	"fmt"
	"testing"
)

const testAll4url, index4url = false, 4

func Test_downloadurl(t *testing.T) {

	songID := map[string]string{
		"netease": "1327566544",
		"tencent": "000bdxcy3ta8Q3",
		"xiami":   "1807124714",
		"kugou":   "d353b69a3b7f1a250000c5daabb8a4f1",
		"baidu":   "578055564"}

	if testAll4url {
		for _, s := range sites {
			resp := Downloadurl(s, songID[s])
			fmt.Println(resp)
		}
	} else {
		resp := Downloadurl(sites[index4url], songID[sites[index4url]])
		fmt.Println(resp)
	}

}

/*
TODO:
	网易云音乐无法获取到下载直链，提示参数错误

    百度音乐不定时出现参数错误？
*/
