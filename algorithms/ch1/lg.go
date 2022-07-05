package ch1

import (
	"math"
	"strings"
)

// 接受一个整型参数N，返回不大于log2N 的最大整数
func minorLg(N int) int {
	if N < 2 {
		return 0
	} else if N == 2 {
		return 1
	}
	return 1 + minorLg(N/2)
}

// 计算ln(N!) 的值
func lnN(N int) float64 {
	if N <= 1 {
		return math.Log(float64(N))
	}
	return math.Log(float64(N)) * math.Log(float64(N-1))
}

type Ordered[T any] interface {
	~int | ~string
}

func min[T Ordered[T]](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Ordered[T]](a, b T) T {
	if a < b {
		return b
	}
	return a
}

// 使用欧几里得法，求最大公约数
func gcd(a, b int32) int32 {
	for {
		r := a % b
		if r == 0 {
			return b
		}
		a, b = b, r
	}
}

// 回环变位判断
func circularRotation(s, t string) bool {
	return strings.Contains(s+s, t)
}
