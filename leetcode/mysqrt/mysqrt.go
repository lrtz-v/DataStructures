package mysqrt

import "math"

/**
  牛顿法求平方根
  X(n+1) = (Xn+a/Xn)/2
*/

func newtonRaphson(x int, err float64) float64 {
	a := float64(x)/2
	for {
		tmp := (a + float64(x) / a) / 2
		if math.Abs(a-tmp) < err {
			return tmp
		}
		a = tmp
	}
}

/**
  二分法
*/
func midSearch(x int) int {
    l, r := 0, x
    ans := -1
    for l <= r {
        mid := l + (r - l) >> 1
        if mid * mid <= x {
            ans = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return ans
}
