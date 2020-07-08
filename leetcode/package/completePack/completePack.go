package completePack

/* 完全背包问题

有 N 种物品和一个容量为 V 的背包，每种物品都有无限件可用。放入第 i 种物品
的费用是 Ci，价值是 Wi。求解：将哪些物品装入背包，可使这些物品的耗费的费用总
和不超过背包容量，且价值总和最大。

F[i, v] = max(F[i − 1, v], F[i, v − Ci] + Wi)
*/
func completePack(N, C int, value, weight []int) int {
	m := make([]int, C+1)
	for i := 1; i < N+1; i++ {
		for w := weight[i-1]; w < C+1; w++ {
			m[w] = max(m[w], m[w - weight[i-1]] + value[i-1])
		}
	}

	return m[C]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
