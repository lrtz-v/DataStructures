package waysToStep

func waysToStep(n int) int {

    a, b, c := 1, 2, 4
    if n <= 2 {
        return n
    }

    for i := 4; i < n+1; i++ {
		a, b, c = b, c, (a+b+c)%1000000007
    }
    return c
}