package musicspider

import (
	"fmt"
	"testing"
)

const testAll4search, index4search = true, 0

func Test_search(t *testing.T) {
	if testAll4search {
		for _, s := range sites {
			resp := Search(s, "hello", map[string]int{"limit": 10})
			fmt.Println(resp)
		}
	} else {
		resp := Search(sites[index4search], "hello", map[string]int{"limit": 10})
		fmt.Println(resp)
	}
}
