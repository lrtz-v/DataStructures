package leetcode

/*
	5座金矿，储量不同，需要人力也不同；10个工人；每座金矿需要人力如下：
	500KG  5人
	400KG  5人
	200KG  3人
	300KG  4人
	350KG  3人

	每座金矿要么全挖，要么不挖，如何获取更多的黄金
*/

/*
workerNum: 工人数量
costList：金矿开采所需人工数量
goldsList：金矿储量
*/
func getBestGildMining(workerNum int, costList, goldsList []int) int {
	resultTable := make([][]int, len(goldsList)+1)
	for i := 0; i < len(goldsList)+1; i++ {
		resultTable[i] = make([]int, workerNum+1)
	}

	for i := 1; i <= len(goldsList); i++ {
		for j := 1; j <= workerNum; j++ {
			if j < costList[i-1] {
				resultTable[i][j] = resultTable[i-1][j]
			} else {
				resultTable[i][j] = max(resultTable[i-1][j], resultTable[i-1][j-costList[i-1]]+goldsList[i-1])
			}
		}
	}
	showTable(resultTable)
	return resultTable[len(goldsList)][workerNum]
}

func getBestGildMining2(workerNum int, costList, goldsList []int) int {
	result := make([]int, workerNum+1)

	for i := 1; i <= len(goldsList); i++ {
		for j := workerNum; j >= 1; j-- {
			if j >= costList[i-1] {
				result[j] = max(result[j], result[j-costList[i-1]]+goldsList[i-1])
			}
			// fmt.Println(result)
		}
	}
	return result[workerNum]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func showTable(table [][]int) {
	rows := len(table)
	cols := len(table[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// fmt.Printf("%d\t", table[i][j])
		}
		// fmt.Println()
	}
}
