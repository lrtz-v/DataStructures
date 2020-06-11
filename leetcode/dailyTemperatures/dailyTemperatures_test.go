package dailyTemperatures

import "testing"

func TestDailyTemperatures(t *testing.T) {

	d := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	a := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if len(d) != len(a) {
		t.Fatal(d)
		t.Fatal("[*] DailyTemparatures Error!")
	}
	for i, v := range d {
		if v != a[i] {
			t.Fatal(d)
			t.Fatal("[*] DailyTemparatures Error!")
		}
	}
}
