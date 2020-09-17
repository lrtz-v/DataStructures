package restoreIpAddresses

import "testing"


func appendElement(res []int, num int) {
	res = append(res, num)
}

func TestSliceCopy(t *testing.T) {
	res := []int{1, 2, 3}

	res = append(res, 4)
	appendElement(res, 5)
	appendElement(res, 6)
	if len(res) != 4 {
		t.Fatal("...")
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	ipList := restoreIpAddresses("25525511135")
	if len(ipList) != 2 {
		
		t.Fatal(ipList)
	}
}
