package bigNumSum

import (
	"math"
	"strconv"
	"strings"
)

/*
	两个超大数字的相加

	426709752318
     95481253129
------------------
    522191005447

*/

// 把每个元素才分出来，单独运算
func bigNumSum(bigNumA, bigNumB string) string {
	if len(bigNumA) == 0 || len(bigNumA) == 0 {
		return ""
	}
	if len(bigNumA) < len(bigNumB) {
		bigNumA, bigNumB = bigNumB, bigNumA
	}
	d := len(bigNumA) - len(bigNumB)
	tmp := make([]int, len(bigNumA)+1)

	for i := len(bigNumA) - 1; i >= 0; i-- {
		numA, _ := strconv.Atoi(string(bigNumA[i]))
		if i < d {
			tmp[i+1] = numA
			continue
		}
		numB, _ := strconv.Atoi(string(bigNumB[i-d]))
		tmp[i+1] = numA + numB
	}

	res := make([]string, len(tmp))
	for i := len(tmp) - 1; i > 0; i-- {
		if tmp[i] < 10 {
			res[i] = strconv.Itoa(tmp[i])
			continue
		}
		val := tmp[i]
		tmp[i] = val % 10
		tmp[i-1] += val / 10
		res[i] = strconv.Itoa(tmp[i])
	}
	return strings.Join(res, "")
}

// 才分成基础数据可以支持运算的大小
func bigNumSum2(bigNumA, bigNumB string, bit int) string {
	if len(bigNumA) == 0 || len(bigNumA) == 0 {
		return ""
	}
	if len(bigNumA) < len(bigNumB) {
		bigNumA, bigNumB = bigNumB, bigNumA
	}
	d := len(bigNumA) - len(bigNumB)
	tmp := make([]int, 1+len(bigNumA)/bit)
	res := make([]string, len(tmp))

	index := len(res) - 1
	for i := len(bigNumA); i >= 0; i -= bit {
		AEndIndex := i
		AStartIndex := i - bit
		BEndIndex := i - d
		BStartIndex := i - d - bit
		if AStartIndex < 0 {
			AStartIndex = 0
		}
		if BStartIndex < 0 {
			BStartIndex = 0
		}

		numA, _ := strconv.Atoi(string(bigNumA[AStartIndex:AEndIndex]))
		if BEndIndex < 0 {
			tmp[index] = numA
			index--
			continue
		}
		numB, _ := strconv.Atoi(string(bigNumB[BStartIndex:BEndIndex]))
		tmp[index] = numA + numB
		index--
	}

	checkNum := int(math.Pow10(bit))
	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] == 0 && i == 0 {
			continue
		}
		if tmp[i] < checkNum {
			res[i] = strconv.Itoa(tmp[i])
			continue
		}
		val := tmp[i]
		tmp[i] = val % checkNum
		tmp[i-1] += val / checkNum
		if tmp[i] == 0 {
			zeroList := make([]string, bit)
			for j := 0; j < bit; j++ {
				zeroList[j] = "0"
			}
			res[i] = strings.Join(zeroList, "")
		} else {
			res[i] = strconv.Itoa(tmp[i])
		}
	}
	return strings.Join(res, "")
}
