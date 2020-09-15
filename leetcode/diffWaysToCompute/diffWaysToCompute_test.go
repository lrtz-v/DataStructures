package diffWaysToCompute

import "testing"

func TestDiffWaysToCompute(t *testing.T) {

	res := diffWaysToCompute("2-1-1")
	t.Fatal(res)
	if len(res) != 2 {
		t.Fatal(res)
	}
}
