package ch1

import "testing"

func TestMinorLg(t *testing.T) {

	test_data := [][]int{
		{2, 1},
		{3, 1},
		{8, 3},
		{16, 4},
		{4, 2},
		{9, 3},
		{25, 4},
	}

	for _, val := range test_data {
		v := minorLg(val[0])
		if v != val[1] {
			t.Fatalf("expect %d, Got: %d\n", val[1], v)
		}
	}
}

func TestGcd(t *testing.T) {
	a, b := 1997, 615
	g := gcd(int32(a), int32(b))
	if g != 1 {
		t.Fatalf("Expect 1, Got %d\n", g)
	}
	a, b = 1111111, 1234567
	g = gcd(int32(a), int32(b))
	if g != 1 {
		t.Fatalf("Expect 1, Got %d\n", g)
	}
}

func TestCircularRotation(t *testing.T) {
	s1 := "ACTGACG"
	s2 := "TGACGAC"
	if !circularRotation(s1, s2) {
		t.Fatal("check circularRotation failed")
	}
}
