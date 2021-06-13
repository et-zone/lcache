package lcache

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	cache := NewLocalCache()
	cache.Set("aa", "aaaa", 0)
	cache.Set("bb", "bbbb", 0)
	cache.Set("cc", "cccc", 0)

	d, _ := cache.Get("aa")
	fmt.Println(d)
	d, _ = cache.Get("bb")
	fmt.Println(d)
	d, _ = cache.Get("cc")
	fmt.Println(d)
}
