package calculate

import (
	"testing"
)

func TestGetNextNum(t *testing.T) {
	num, i := getNextNum("123 ", 0)
	if num != 123 || i != 4 {
		t.Fatal("TestGetNextNum Failed.")
	}
}

func TestCalculate(t *testing.T) {
	num := calculate("3+2*2")
	if num != 7 {
		t.Fatal("TestCalculate Failed.")
	}
}

func TestCalculate1(t *testing.T) {
	num := calculate(" 3/2 ")
	if num != 1 {
		t.Fatal("TestCalculate Failed.")
	}
}

func TestCalculate2(t *testing.T) {
	num := calculate(" 3+5 / 2 ")
	if num != 5 {
		t.Fatal("TestCalculate Failed.")
	}
}

func TestCalculate3(t *testing.T) {
	num := calculate("0")
	if num != 0 {
		t.Fatal("TestCalculate Failed.")
	}
}

func TestCalculate4(t *testing.T) {
	num := calculate("1+1+1")
	if num != 3 {
		t.Fatal("TestCalculate Failed.")
	}
}

func TestCalculate5(t *testing.T) {
	num := calculate("2*3+4")
	if num != 10 {
		t.Fatal("TestCalculate Failed.")
	}
}
