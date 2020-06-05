package greatestCommonDivisor

func gcd(a, b int) int {
	if a == b {
		return a
	}

	// 当a、b均为偶数时，gcd(a,b) = 2* gcd(a/2, b/2)
	if a&1 == 0 && b&1 == 0 {
		return gcd(a>>1, b>>1) << 1
	} else if a&1 == 0 && b&1 == 1 { // 当a是偶数，b是奇数，gcd(a,b) = gcd(a/2, b)
		return gcd(a>>1, b)
	} else if a&1 == 1 && b&1 == 0 { // danga是奇数，b是偶数，gcd(a,b) = gcd(a, b/2)
		return gcd(a, b>>1)
	}
	// a,b均为奇数 gcd(a, b) = gcd(b, a-b)
	big, small := a, b
	if b > a {
		big, small = b, a
	}
	return gcd(big-small, small)
}
