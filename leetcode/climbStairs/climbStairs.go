package climbStairs

var c = make(map[int]int)

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    first := 1
    second := 2
    third := 0
    for i:=3;i<=n;i++ {
        third = first+second
        second, first = third, second
    }
    return third
}
