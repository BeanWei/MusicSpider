package musicspider

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	for i, s := range sites {
		if i == 1 {
			resp := search("hello", s, map[string]int{"limit": 10})
			fmt.Println(resp)
		}
	}

}
