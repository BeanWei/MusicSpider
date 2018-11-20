package musicspider

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	resp := search("hello", "netease", map[string]int{"limit": 10})
	fmt.Println(resp)
}
