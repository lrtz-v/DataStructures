package bitMap

import (
	"testing"
)

func TestBitMap(t *testing.T) {
	t.Skip()
	var bitmap32 Bitmap32
	bitmap32 = bitmap32.SetBit(10)
	bitmap32 = bitmap32.SetBit(7)
	// fmt.Println(bitmap32)
	if !bitmap32.GetBit(10) {
		t.Error("Expected true, Got: false")
	}
	if !bitmap32.GetBit(7) {
		t.Error("Expected true, Got: false")
	}
}
