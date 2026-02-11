package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 10 * time.Second
	cache := NewCache(interval)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("example.com-1"),
		},
		{
			key: "12345",
			val: []byte("12345"),
		},
		{
			key: "https://example.com",
			val: []byte("example.com-2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("running case num: %v", i), func(t *testing.T) {
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("cannot find a key")
				return
			} else if string(val) != string(c.val) {
				t.Errorf("values are wrong")
				return
			}
		})
	}
}

func TestCacheReapLoop(t *testing.T) {
	const interval = 5 * time.Second
	const waitTime = interval + 10 * time.Millisecond

	cache := NewCache(interval)
	cache.Add("example.com", []byte("testdata"))

	_, ok := cache.Get("example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}