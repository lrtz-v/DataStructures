package mypow

func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickMul(x, n)
    }
    return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
    ans := 1.0
    tmp := x
    for N > 0 {
        if N % 2 == 1 {
            ans *= tmp
        }
        tmp *= tmp
        N /= 2
    }
    return ans
}
