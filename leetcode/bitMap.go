package leetcode

// Bitmap32 tracks 32 bool values within a uint32
type Bitmap32 uint32

func (b Bitmap32) SetBit(pos uint) Bitmap32 {
	return b | (1 << pos)
}

func (b Bitmap32) GetBit(pos uint) bool {
	return (b & (1 << pos)) != 0
}
