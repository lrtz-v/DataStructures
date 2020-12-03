package entry

import (
	"reflect"
	"testing"
	"time"
)

func TestEntry(t *testing.T) {
	k := Key("123")

	if reflect.TypeOf(k).Name() != "Key" {
		t.Fatalf("Get type error: %s Expected Key\n", reflect.TypeOf(k).Name())
	}

	kstr := string(k)
	if reflect.TypeOf(kstr).Name() != "string" {
		t.Fatalf("Get type error: %s Expected string\n", reflect.TypeOf(k).Name())
	}
}

func TestHash(t *testing.T) {
	seed := uint64(time.Now().Nanosecond())
	bucketmask := uintptr(128 - 1)

	count := make([]int, 128)

	for i := 0; i < 10000; i++ {
		k := Key(string(time.Now().Format("Jan _2 15:04:05.000000000")))
		hash := k.Hash(seed)		
		bucket := hash & bucketmask
		count[bucket]++
	}
}
