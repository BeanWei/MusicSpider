package musicspider

import (
	"fmt"
	"testing"
)

const testAll, index = false, 1

func Test_search(t *testing.T) {
	if testAll {
		for _, s := range sites {
			resp := search("hello", s, map[string]int{"limit": 10})
			fmt.Println(resp)
		}
	} else {
		resp := search("hello", sites[index], map[string]int{"limit": 10})
		fmt.Println(resp)
	}

}
