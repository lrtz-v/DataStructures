package LRUCache

import (
	"testing"
)

func TestLRUcache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	v := cache.Get(1) // 返回  1
	if v != 1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 1 Expected 1 Got %d\n", v)
	}
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	v = cache.Get(2) // 返回 -1 (未找到)
	if v != -1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 2 Expected -1 Got %d\n", v)
	}
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	v = cache.Get(1)    // 返回 -1 (未找到)
	if v != -1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 1 Expected -1 Got %d\n", v)
	}
	v = cache.Get(3) // 返回  3
	if v != 3 {
		t.Fatalf("[*] LRUcache ERROR! Key: 3 Expected 3 Got %d\n", v)
	}
	v = cache.Get(4) // 返回  4
	if v != 4 {
		t.Fatalf("[*] LRUcache ERROR! Key: 4Expected 4 Got %d\n", v)
	}
}

func TestLRUcache2(t *testing.T) {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.Put(2, 2)
	v := cache.Get(2)
	if v != 2 {
		t.Fatalf("[*] LRUcache ERROR! Key: 2 Expected 2 Got %d\n", v)
	}
	cache.Put(1, 1)
	cache.Put(4, 1)
	v = cache.Get(2)
	if v != -1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 2 Expected -1 Got %d\n", v)
	}
}

func TestLRUcache3(t *testing.T) {
	cache := Constructor(1)

	cache.Put(2, 1)
	v := cache.Get(2)
	if v != 1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 2 Expected 1 Got %d\n", v)
	}
	cache.Put(3, 2)
	v = cache.Get(2)
	if v != -1 {
		t.Fatalf("[*] LRUcache ERROR! Key: 2 Expected -1 Got %d\n", v)
	}
	v = cache.Get(3)
	if v != 2 {
		t.Fatalf("[*] LRUcache ERROR! Key: 3 Expected 2 Got %d\n", v)
	}
}

func TestLRUcache4(t *testing.T) {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.Put(3, 2)
	cache.Get(3)
	cache.Get(2)
	cache.Put(4, 3)
	cache.Get(2)
	cache.Get(3)
	cache.Get(4)
}
