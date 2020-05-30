package largestRectangleArea
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

*/

func largestRectangleArea(heights []int) int {
    n := len(heights)
    left, right := make([]int, n), make([]int, n)
    monoStack := []int{}
    for i := 0; i < n; i++ {
        for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
            monoStack = monoStack[:len(monoStack)-1]
        }
        if len(monoStack) == 0 {
            left[i] = -1
        } else {
            left[i] = monoStack[len(monoStack)-1]
        }
        monoStack = append(monoStack, i)
    }
    monoStack = []int{}
    for i := n - 1; i >= 0; i-- {
        for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
            monoStack = monoStack[:len(monoStack)-1]
        }
        if len(monoStack) == 0 {
            right[i] = n
        } else {
            right[i] = monoStack[len(monoStack)-1]
        }
        monoStack = append(monoStack, i)
    }
    ans := 0
    for i := 0; i < n; i++ {
        ans = max(ans, (right[i] - left[i] - 1) * heights[i])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func largestRectangleArea2(heights []int) int {
    n := len(heights)
    left, right := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        right[i] = n
    }
    monoStack := []int{}
    for i := 0; i < n; i++ {
        for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
            right[monoStack[len(monoStack)-1]] = i
            monoStack = monoStack[:len(monoStack)-1]
        }
        if len(monoStack) == 0 {
            left[i] = -1
        } else {
            left[i] = monoStack[len(monoStack)-1]
        }
        monoStack = append(monoStack, i)
    }
    ans := 0
    for i := 0; i < n; i++ {
        ans = max(ans, (right[i] - left[i] - 1) * heights[i])
    }
    return ans
}
