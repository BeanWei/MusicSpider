package musicspider

import (
	"fmt"
	"testing"
)

const testAll4url, index4url = false, 2

func Test_downloadurl(t *testing.T) {

	songID := map[string]string{
		"netease": "35847388",
		"tencent": "000bdxcy3ta8Q3",
		"xiami":   "1807124714",
		"kugou":   "d353b69a3b7f1a250000c5daabb8a4f1",
		"baidu":   "578055564"}

	if testAll4url {
		for _, s := range sites {
			resp := downloadurl(s, songID[s])
			fmt.Println(resp)
		}
	} else {
		resp := downloadurl(sites[index4url], songID[sites[index4url]])
		fmt.Println(resp)
	}

}

/*
Note:
网易云音乐无法获取到下载直链，提示参数错误

QQ音乐获取到的直链打开提示不被允许

虾米音乐 https://m128.xiami.net/204/7204/2104245668/1807124714_1542698603227.mp3?auth_key=1543374000-0-0-5f5025aeb43cf1a2d4ee6dc967d0f191
		http://m128.xiami.net/204/7204/2104245668/1807124714_1542698603227.mp3?auth_key=1543374000-0-0-5f5025aeb43cf1a2d4ee6dc967d0f191
		https://m128.xiami.net/202/2202/2104221786/1806947021_1542783344343.mp3?auth_key=1543374000-0-0-7601d9f243f5aecdf3b357649bd42ac8
*/
