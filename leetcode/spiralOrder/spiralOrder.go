package spiralOrder


/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [
	[1,2,3],
	[4,5,6],
	[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [
	[1, 2, 3, 4],
	[5, 6, 7, 8],
	[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7] 

限制：
0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

*/
func spiralOrder(matrix [][]int) []int {

	res := []int{}
    if len(matrix) == 0 {
        return res   
    }

	l := 0  // 左
	r := len(matrix[0]) - 1  // 右
	t := 0  // 上
	b := len(matrix)-1  // 下

	for {
		// 左➡️右
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		// 上⬇️下
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		// 右⬅️左
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if t > b {
			break
		}

		// 下⬆️上
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}

	}


	return res
}