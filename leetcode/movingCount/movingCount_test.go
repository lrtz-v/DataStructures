package movingcount

import "testing"

func TestMovingCount(t *testing.T) {
	count := movingCount(2, 3, 1)
	if count != 3 {
		t.Fatal("[*] TestMovingCount failed.")
	}

	count = movingCount(3, 1, 0)
	if count != 1 {
		t.Fatal("[*] TestMovingCount failed.")
	}
}

func TestMovingCount2(t *testing.T) {
	count := movingCount(16, 8, 4)
	if count != 15 {
		t.Fatalf("\n[*] TestMovingCount failed. Expected 15 Got %d\n", count)
	}
}

func TestCalculate(t *testing.T) {
	s := calculate(100)
	if s != 1 {
		t.Fatal("[*] TestCalculate Failed.")
	}

	s = calculate(101)
	if s != 2 {
		t.Fatal("[*] TestCalculate Failed.")
	}

	s = calculate(99)
	if s != 18 {
		t.Fatal("[*] TestCalculate Failed.")
	}
}

func TestArrayEqual(t *testing.T) {
	a := [1]int{1}
	b := [1]int{1}
	if a != b {
		t.Fatal("[*] TestArrayEqual .1 Failed")
	}

	c := make([]int, 1)
	c = append(c, 1)
	d := make([]int, 1)
	d = append(d, 1)

	// != is not for slice
	// if c != d {
	// 	t.Fatal("[*] TestArrayEqual .2 Failed")
	// }
}
