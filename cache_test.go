package lcache

import (
	"fmt"
	"testing"
	// "time"
)

func Test_run(t *testing.T) {
	cache := NewLocalCache()
	cache.Set("aa", "aaaa", 2)
	cache.Set("bb", "bbbb", 0)
	cache.Set("cc", "cccc", 0)

	d, _ := cache.Get("aa")
	fmt.Println(d)
	d, _ = cache.Get("bb")
	fmt.Println(d)
	d, _ = cache.Get("cc")
	fmt.Println(d)

	// time.Sleep(time.Secosnd * 3)
	// d, _ = cache.Get("aa")
	// fmt.Println(d)

}
