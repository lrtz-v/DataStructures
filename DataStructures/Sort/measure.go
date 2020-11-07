package sort

import (
	"math/rand"
	"time"
)

func measure(alg string, l []int64) float64 {
	now := time.Now()

	switch alg {
	case "Insertion":
		insertion := Insertion{}
		insertion.Sort(l)
	case "Selection":
		selection := Selection{}
		selection.Sort(l)
	case "Shell":
		shell := Shell{}
		shell.Sort(l)
	}

	return time.Since(now).Seconds()
}

func randomList(n int) []int64 {
	l := []int64{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		l = append(l, rand.Int63())
	}
	return l
}

func randomSortedList(n int) []int64 {
	l := []int64{}

	for i := 0; i < n; i++ {
		l = append(l, int64(i))
	}

	return l
}
