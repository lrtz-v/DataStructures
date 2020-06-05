package bigNumSum

import (
	"testing"
)

func TestBigNumSum(t *testing.T) {
	t.Skip()
	res := bigNumSum("426709752318", "95481253129")
	if res != "522191005447" {
		t.Errorf("Expected %s, Got: %s\n", "522191005447", res)
	}
}

func TestBigNumSum0(t *testing.T) {
	t.Skip()
	res := bigNumSum("4267097", "1234")
	if res != "4268331" {
		t.Errorf("Expected %s, Got: %s\n", "4268331", res)
	}
}

func TestBigNumSum2(t *testing.T) {
	t.Skip()
	res := bigNumSum2("426709752318", "95481253129", 2)
	if res != "522191005447" {
		t.Errorf("Expected %s, Got: %s\n", "522191005447", res)
	}
}

func TestBigNumSum3(t *testing.T) {
	t.Skip()
	res := bigNumSum2("4267097", "1234", 2)
	if res != "4268331" {
		t.Errorf("Expected %s, Got: %s\n", "4268331", res)
	}
}

func TestBigNumSum4(t *testing.T) {
	t.Skip()
	res := bigNumSum2("426709752318", "95481253129", 2)
	if res != "522191005447" {
		t.Errorf("Expected %s, Got: %s\n", "522191005447", res)
	}
}

func TestBigNumSum5(t *testing.T) {
	t.Skip()
	res := bigNumSum2("4267097", "1234", 3)
	if res != "4268331" {
		t.Errorf("Expected %s, Got: %s\n", "4268331", res)
	}
}

func TestBigNumSum6(t *testing.T) {
	t.Skip()
	res := bigNumSum2("426709752318", "95481253129", 4)
	if res != "522191005447" {
		t.Errorf("Expected %s, Got: %s\n", "522191005447", res)
	}
}

func TestBigNumSum7(t *testing.T) {
	t.Skip()
	res := bigNumSum2("4267097", "1234", 1)
	if res != "4268331" {
		t.Errorf("Expected %s, Got: %s\n", "4268331", res)
	}
}

func TestBigNumSum8(t *testing.T) {
	t.Skip()
	res := bigNumSum2("10000001", "10000001", 1)
	if res != "20000002" {
		t.Errorf("Expected %s, Got: %s\n", "20000002", res)
	}
}
