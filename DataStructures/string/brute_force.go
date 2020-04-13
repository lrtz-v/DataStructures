package string

/**
 * Brute Force 暴力匹配算法(朴素匹配算法)
 * 在字符串A中查找字符串B
 * A: 主串    len：n
 * B: 模式串  len：m
 * n > m
 */

func bruteForce(A, B string) (index int) {

	index = -1
	for i := 0; i < len(A)-len(B); i++ {
		if A[i:i+len(B)] == B {
			index = i
			break
		}
	}

	return index

}
