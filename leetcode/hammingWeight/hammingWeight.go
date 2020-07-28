package hammingWeight

func hammingWeight(num uint32) int {
    if num == 0 {
        return 0
	}
    return int(num % 2) + hammingWeight(num >> 1)
}

/*
若 n & 1 = 0，则 n 二进制 最右一位 为 0
若 n & 1 = 1，则 n 二进制 最右一位 为 1
*/
func hammingWeight2(num uint32) int {
	res := 0
	for ; num > 0; {
		res += int(num & 1)
		num >>= 1
	}
	return res
}

/*
(n−1) 解析： 二进制数字 n 最右边的 1 变成 0 ，此 1 右边的 0 都变成 1
n & (n - 1) 解析： 二进制数字 n 最右边的 1 变成 0 ，其余不变
*/
func hammingWeight3(num uint32) int {
	res := 0
	for ; num > 0; {
		res++
		num &= num - 1	
	}
	return res
}
