package package01

/*
给定一组n个物品，每种物品都有自己的重量（Wi）和价值（Vi），
在限定的总重量/总容量（C）内，选择其中若干个（也即每种物品可以选0个或1个），设计选择方案使得物品的总价值最高。

解题：

定义子问题 P(i, W) 为：在前 i 个物品中挑选总重量不超过 W 的物品，每种物品至多只能挑选1个，使得总价值最大
这时的最优值记作 m(i, W) ，其中 i <= i <= n, 1 <= W <= C.

考虑第 i 个物品，无外乎两种可能：选，或者不选。
1. 不选的话，背包的容量不变，改变为问题 P(i-1, W)
2. 选的话，背包的容量变小，改变为问题 P(i-1, W-Wi)

m(i, W) = max(m(i-1, W), m(i-1, W-Wi)+Vi)

*/

/*
N: num of items
C: max weight
value: value[i] means value of items
weight: weight[i] means weight
len(value) == N == len(weight)
*/
func package01(N, C int, value, weight []int) int {

	m := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		m[i] = make([]int, C+1)
	}
	for i := 0; i < C+1; i++ {
		m[0][i] = 0
	}

	for i := 1; i < N+1; i++ {
		m[i][0] = 0
	}

	for i := 1; i < N+1; i++ {
		for w := 1; w < C+1; w++ {
			if weight[i-1] > w {
				m[i][w] = m[i-1][w]
			} else {
				m[i][w] = max(m[i-1][w], value[i-1] + m[i-1][w - weight[i-1]])
			}
		}
	}
	return m[N][C]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 空间复杂度优化，使用一维数组代替二维数组
func zeroOnePack(N, C int, value, weight []int) int {
	m := make([]int, C+1)
	for i := 1; i < N+1; i++ {
		for w := C; w > 0; w-- {
			if weight[i-1] <= w {
				m[w] = max(m[w], value[i-1] + m[w - weight[i-1]])
			}
		}
	}

	return m[C]
}