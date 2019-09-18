package concurrency

import (
	"fmt"
	"sort"
	"testing"
)

func TestParallelism(t *testing.T) {
	m, err := ParallelismMD5All(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
